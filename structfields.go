package json

import (
	"reflect"
	"sync"
)

var (
	constantizerType = reflect.TypeOf((*Constantizer)(nil)).Elem()
	omitAlwaysType   = reflect.TypeOf((*OmitAlways)(nil)).Elem()
)

// structField holds pre-computed information about a struct field
// relevant to unmarshaling.
type structField struct {
	index        int
	indexPath    []int    // full path of field indices through embedded structs
	name         string   // JSON key name (from tag or Go field name)
	skip         bool     // json:"-"
	omitempty    bool
	nullempty    bool
	modifiers    []string
	fieldType    reflect.Type
	isConst      bool   // implements Constantizer
	constTag     string // value of json.value tag (for Const[T])
	isOmitAlways bool   // implements OmitAlways
}

// fieldByIndexPath traverses a reflect.Value using the index path to reach the target field.
func fieldByIndexPath(v reflect.Value, path []int) reflect.Value {
	for _, idx := range path {
		v = v.Field(idx)
	}
	return v
}

// structInfo holds cached information about a struct type's fields.
type structInfo struct {
	fields    []structField
	nameIndex map[string]int // JSON name -> index in fields slice
}

var structInfoCache sync.Map // map[reflect.Type]*structInfo

// getStructInfo returns cached struct field info, computing it on first access.
func getStructInfo(t reflect.Type) *structInfo {
	if cached, ok := structInfoCache.Load(t); ok {
		return cached.(*structInfo)
	}

	info := buildStructInfo(t)
	actual, _ := structInfoCache.LoadOrStore(t, info)
	return actual.(*structInfo)
}

// buildStructInfo computes struct field information from a reflect.Type.
func buildStructInfo(t reflect.Type) *structInfo {
	info := &structInfo{
		fields:    make([]structField, 0, t.NumField()),
		nameIndex: make(map[string]int, t.NumField()),
	}

	collectFields(t, info, nil)

	return info
}

// collectFields recursively collects struct fields, flattening anonymous (embedded) struct fields.
// indexPath tracks the chain of field indices needed to reach a field through nested embeddings.
func collectFields(t reflect.Type, info *structInfo, indexPath []int) {
	for i := 0; i < t.NumField(); i++ {
		rf := t.Field(i)
		if !rf.IsExported() {
			continue
		}

		// If the field is an anonymous (embedded) struct, recurse into it.
		if rf.Anonymous && rf.Type.Kind() == reflect.Struct {
			// Check if it implements OmitAlways — if so, skip.
			if rf.Type.Implements(omitAlwaysType) || reflect.PointerTo(rf.Type).Implements(omitAlwaysType) {
				continue
			}

			collectFields(rf.Type, info, append(append([]int(nil), indexPath...), i))
			continue
		}

		fullPath := append(append([]int(nil), indexPath...), i)

		sf := structField{
			index:     i,
			indexPath: fullPath,
			name:      rf.Name,
			fieldType: rf.Type,
		}

		// Check if the field type implements Constantizer or OmitAlways.
		if rf.Type.Implements(constantizerType) || reflect.PointerTo(rf.Type).Implements(constantizerType) {
			sf.isConst = true
			if tag, found := rf.Tag.Lookup("json.value"); found {
				sf.constTag = tag
			}
		}
		if rf.Type.Implements(omitAlwaysType) || reflect.PointerTo(rf.Type).Implements(omitAlwaysType) {
			sf.isOmitAlways = true
		}

		// Parse the json struct tag.
		tag, found := rf.Tag.Lookup("json")
		if found {
			var newname string
			newname, sf.skip, sf.omitempty, sf.nullempty, sf.modifiers = parseTag(tag)
			if "" != newname {
				sf.name = newname
			}
		}

		if sf.skip {
			continue
		}

		idx := len(info.fields)
		info.fields = append(info.fields, sf)
		info.nameIndex[sf.name] = idx
	}
}
