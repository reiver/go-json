package json

import (
	"encoding"
	gojson "encoding/json"
	"fmt"
	"reflect"
	"sort"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-lck"
)

// Usher marshals a Go type into JSON.
//
// If you want the "string" modifier you will need to call json.Usher.ImplantModifier() to add it:
//
//	var jsonUsher json.Usher
//
//	jsonUsher.ImplantModifier("string", json.StringModifierFunc)
//
// You can also add your own modifiers:
//
//	var jsonUsher json.Usher
//
//	jsonUsher.ImplantModifier("digest", digestFunc)
type Usher struct {
	modifierFuncs lck.Lockable[map[string]ModifierFunc]
}

func (receiver *Usher) ImplantModifier(name string, fn ModifierFunc) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.modifierFuncs.Let(func(modifiers *map[string]ModifierFunc){
		if nil == *modifiers {
			*modifiers = map[string]ModifierFunc{}
		}

		switch fn {
		case nil:
			delete(*modifiers, name)
		default:
			(*modifiers)[name] = fn
		}
	})
}

func (receiver *Usher) marshalMap(value any) ([]byte, error) {
	if nil == value {
		return []byte{'n','u','l','l'}, nil
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '{')

	{
		var reflectedValue = reflect.ValueOf(value)

		var reflectedKeys []reflect.Value = reflectedValue.MapKeys()

		{
			var fn = func(index1, index2 int) bool {
				value1 := reflectedKeys[index1]
				value2 := reflectedKeys[index2]

				return value1.String() < value2.String()
			}
			sort.Slice(reflectedKeys, fn)
		}

		for i, reflectedKey := range reflectedKeys {

			mapValueAny := reflectedValue.MapIndex(reflectedKey).Interface()

			switch mapValueAny.(type) {
			case OmitAlways:
				continue
			}

			if 0 < i {
				p = append(p, ',')
			}

			{
				keyAny := reflectedKey.Interface()

				var encoded []byte
				var err error
				switch casted := keyAny.(type) {
				case encoding.TextMarshaler:
					var bytes []byte
					bytes, err = casted.MarshalText()
					if nil != err {
						return nil, erorr.Errorf("json: problem text-marshaling key of type %T (which also is a text-marshaler): %w", keyAny, err)
					}
					encoded, err = receiver.Marshal(string(bytes))
					if nil != err {
						return nil, erorr.Errorf("json: problem json-marshaling key of type %T (which also is a text-marshaler): %w", keyAny, err)
					}
				case fmt.Stringer:
					encoded, err = receiver.Marshal(casted.String())
					if nil != err {
						return nil, erorr.Errorf("json: problem json-marshaling key of type %T (which also is a stringer): %w", keyAny, err)
					}
				case string:
					encoded, err = receiver.Marshal(casted)
					if nil != err {
						return nil, erorr.Errorf("json: problem json-marshaling key of type %T (string): %w", keyAny, err)
					}
				case []byte:
					encoded, err = receiver.Marshal(string(casted))
					if nil != err {
						return nil, erorr.Errorf("json: problem json-marshaling key of type %T ([]byte): %w", keyAny, err)
					}
				case []rune:
					encoded, err = receiver.Marshal(string(casted))
					if nil != err {
						return nil, erorr.Errorf("json: problem json-marshaling key of type %T ([]rune): %w", keyAny, err)
					}
				default:
					return nil, erorr.Errorf("json: cannot json-marshal a key of type %T", keyAny)
				}
				p = append(p, encoded...)
				p = append(p, ':')
			}

			{
				encoded, err := receiver.Marshal(mapValueAny)
				if nil != err {
					return nil, erorr.Errorf("json: cannot json-marshal a map-value of type %T", mapValueAny)
				}
				p = append(p, encoded...)
			}
		}
	}

	p = append(p, '}')

	return p, nil
}

func (receiver *Usher) marshalSlice(value any) ([]byte, error) {
	if nil == value {
		return []byte{'n','u','l','l'}, nil
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '[')

	var reflectedValue = reflect.ValueOf(value)
	var length int =  reflectedValue.Len()

	var notempty bool

	for i:=0; i<length; i++ {
		elementReflectedValue := reflectedValue.Index(i)

		if notempty {
			p = append(p, ',')
		}
		{
			var element any = elementReflectedValue.Interface()

			bytes, err := receiver.Marshal(element)
			if nil != err {
				return nil, erorr.Errorf("json: problem marshaling element %d of the slice %T: %w", i, value, err)
			}

			p = append(p, bytes...)
		}
		notempty = true
	}

	p = append(p, ']')

	return p, nil
}


func (receiver *Usher) marshalStruct(value any) ([]byte, error) {
	if nil == value {
		return nil, errNotStruct
	}

	reflectedType := reflect.TypeOf(value)
	if nil == reflectedType {
		return nil, errBadReflection
	}

	if reflect.Struct != reflectedType.Kind() {
		return nil, errNotStruct
	}

	reflectedValue := reflect.ValueOf(value)

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '{')

	{
		var notempty bool

		var numfields int = reflectedType.NumField()

		for i:=0; i<numfields; i++ {
			reflectedStructFieldType  := reflectedType.Field(i)
			if !reflectedStructFieldType.IsExported() {
		/////////////// CONTINUE
				continue
			}

			reflectedStructFieldValue := reflectedValue.Field(i)
			reflectedStructFieldValueAny := reflectedStructFieldValue.Interface()

			switch reflectedStructFieldValueAny.(type) {
			case OmitAlways:
				continue
			}

			var name string
			var skip bool
			var omitempty bool
			var modifiers []string
			{
				name = reflectedStructFieldType.Name
			}
			{
				tag, found := reflectedStructFieldType.Tag.Lookup("json")
				if found {
					var newname string

					newname, skip, omitempty, modifiers = parseTag(tag)

					if "" != newname {
						name = newname
					}
				}
			}

			if skip {
		/////////////// CONTINUE
				continue
			}
			if omitempty {
				switch casted := reflectedStructFieldValueAny.(type) {
				case Emptier:
					if casted.IsEmpty() {
						continue
					}
				case Nothinger:
					if casted.IsNothing() {
						continue
					}
				}

				var empty reflect.Value = reflect.Zero(reflectedStructFieldType.Type)

				if reflect.DeepEqual(empty.Interface(), reflectedStructFieldValue.Interface()) {
		/////////////////////// CONTINUE
					continue
				}
			}

			{
				var valuebytes []byte
				{
					var fieldvalue any = reflectedStructFieldValue.Interface()

					var err error
					valuebytes, err = receiver.Marshal(fieldvalue)
					if nil != err {
						if omitempty && erorr.Is(err, ErrEmpty("")) {
		/////////////////////////////////////// CONTINUE
							continue
						}
						return nil, erorr.Errorf("json: problem marshaling %T into JSON: %w", fieldvalue, err)
					}

					for _, modifierName := range modifiers {
						if "" == modifierName {
							continue
						}

						var fn ModifierFunc
						receiver.modifierFuncs.Let(func(m *map[string]ModifierFunc){
							if nil == m {
								return
							}
							if nil == *m {
								return
							}

							fn = (*m)[modifierName]
						})

						if nil != fn {
							valuebytes, err = fn(valuebytes)
							if nil != err {
								if omitempty && erorr.Is(err, ErrEmpty("")) {
		/////////////////////////////////////////////////////// CONTINUE
									continue
								}
								return nil, erorr.Errorf("json: problem marshaling %T into JSON using modifier %q: %w", fieldvalue, modifierName, err)
							}
						}
					}

				}

				var namebytes []byte
				{
					var err error

					namebytes, err = gojson.Marshal(name)
					if nil != err {
						return nil, erorr.Errorf("json: problem marshaling string %q into JSON", name)
					}
				}

				if notempty {
					p = append(p, ',')
				}
				p = append(p, namebytes...)
				p = append(p, ':')
				p = append(p, valuebytes...)
				notempty = true
			}
		}
	}

	p = append(p, '}')

	return p, nil
}

func (receiver *Usher) MergeAndMarshal(values ...any) ([]byte, error) {
	var length int = len(values)

	if length <= 0 {
		var empty = [2]byte{'{','}'}
		return empty[:], nil
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var comma bool = false

	bytes = append(bytes, '{')

	for index, value := range values {
		result, err := receiver.Marshal(value)
		if nil != err {
			return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T): %w", 1+index, length, value, err)
		}

		// remove the beginning '{' and the ending '}'.
		{
			var lengthResult int = len(result)

			if lengthResult < 2 {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it is too short (length=%d) to be JSON objects — it is: %q", 1+index, length, value, lengthResult, result)
			}

			if '{' != result[0] {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it does not start with '{' — it is: %q", 1+index, length, value, result)
			}
			if '}' != result[lengthResult-1] {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it does not end with '}' — it is: %q", 1+index, length, value, result)
			}

			result = result[1:lengthResult-1]

			// was this {}?, if "yes", then skip it
			if len(result) <= 0 {
				continue
			}
		}

		if comma {
			bytes = append(bytes, ',')
		}
		bytes = append(bytes, result...)
		comma = true
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
