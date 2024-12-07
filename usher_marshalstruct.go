package json

import (
	gojson "encoding/json"
	"reflect"

	"github.com/reiver/go-erorr"
)

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

				switch reflectedStructFieldValue.Kind() {
				case reflect.Slice,reflect.Array:
					if reflectedStructFieldValue.Len() <= 0 {
		/////////////////////////////// CONTINUE
						continue
					}
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
