package json

import (
	"encoding"
	gojson "encoding/json"
	"reflect"

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

// Marshal return the JSON version of 'value'.
func (receiver *Usher) Marshal(value any) ([]byte, error) {
	if nil == value {
		return gojson.Marshal(value)
	}

	switch casted := value.(type) {
	case Marshaler:
			return casted.MarshalJSON()
	case encoding.TextMarshaler:
		bytes, err := casted.MarshalText()
		if nil != err {
			return nil, err
		}

		var str string = string(bytes)
		return gojson.Marshal(str)
	default:
		reflectedType := reflect.TypeOf(value)
		if nil == reflectedType {
			return nil, errBadReflection
		}

		switch reflectedType.Kind() {
		case reflect.Struct:
			return receiver.marshalStruct(value)
		default:
			return gojson.Marshal(value)
		}
	}
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
					valuebytes, err = Marshal(fieldvalue)
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
