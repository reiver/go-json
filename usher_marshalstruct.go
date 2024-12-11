package json

import (
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
			structFieldValueAny := reflectedStructFieldValue.Interface()

			switch casted := structFieldValueAny.(type) {
			case OmitAlways:
				continue
			case Constantizer:
				const tagName string = "json.value"

				tag, found := reflectedStructFieldType.Tag.Lookup(tagName)
				if !found {
					continue
				}

				value, err := casted.DecodeFromString(tag)
				if nil != err {
					return nil, erorr.Errorf("json: problem decoding '%s' struct-field tag on struct-field of type %T: %w", tagName, casted, err)
				}

				structFieldValueAny = value
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
				switch casted := structFieldValueAny.(type) {
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

				if reflect.DeepEqual(empty.Interface(), structFieldValueAny) {
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
					var err error
					valuebytes, err = receiver.Marshal(structFieldValueAny)
					if nil != err {
						if _, isErrorEmpty := err.(ErrorEmpty) ; omitempty && isErrorEmpty {
		/////////////////////////////////////// CONTINUE
							continue
						}
						return nil, erorr.Errorf("json: [1] problem marshaling %T into JSON: %w", structFieldValueAny, err)
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
								if _, isErrorEmpty := err.(ErrorEmpty) ; omitempty && isErrorEmpty {
		/////////////////////////////////////////////////////// CONTINUE
									continue
								}
								return nil, erorr.Errorf("json: [2] problem marshaling %T into JSON using modifier %q: %w", structFieldValueAny, modifierName, err)
							}
						}
					}

				}

				var namebytes []byte = MarshalString(name)

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
