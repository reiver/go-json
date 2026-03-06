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
	numFields := t.NumField()

	info := &structInfo{
		fields:    make([]structField, 0, numFields),
		nameIndex: make(map[string]int, numFields),
	}

	for i := 0; i < numFields; i++ {
		rf := t.Field(i)
		if !rf.IsExported() {
			continue
		}

		sf := structField{
			index:     i,
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

	return info
}
