package json

import (
	"reflect"

	"codeberg.org/reiver/go-erorr"
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

	var err error
	p, _, err = receiver.marshalStructFields(reflectedType, reflectedValue, p, false)
	if nil != err {
		return nil, err
	}

	p = append(p, '}')

	return p, nil
}

// marshalStructFields iterates over the fields of a struct, marshaling each one.
// For anonymous (embedded) struct fields, it recurses into them to flatten their fields.
func (receiver *Usher) marshalStructFields(reflectedType reflect.Type, reflectedValue reflect.Value, p []byte, notempty bool) ([]byte, bool, error) {

	var numfields int = reflectedType.NumField()

	for i:=0; i<numfields; i++ {
		reflectedStructFieldType  := reflectedType.Field(i)
		if !reflectedStructFieldType.IsExported() {
	/////////////// CONTINUE
			continue
		}

		reflectedStructFieldValue := reflectedValue.Field(i)

		// If this is an anonymous (embedded) struct field, recurse into it.
		if reflectedStructFieldType.Anonymous && reflectedStructFieldType.Type.Kind() == reflect.Struct {
			// Check if it implements OmitAlways — if so, skip.
			if _, ok := reflectedStructFieldValue.Interface().(OmitAlways); ok {
				continue
			}
			var err error
			p, notempty, err = receiver.marshalStructFields(reflectedStructFieldType.Type, reflectedStructFieldValue, p, notempty)
			if nil != err {
				return nil, false, err
			}
			continue
		}

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
				return nil, false, erorr.Errorf("json: problem decoding '%s' struct-field tag on struct-field of type %T: %w", tagName, casted, err)
			}

			structFieldValueAny = value
		}

		var name string
		var skip bool
		var omitempty bool
		var nullempty bool
		var modifiers []string
		{
			name = reflectedStructFieldType.Name
		}
		{
			tag, found := reflectedStructFieldType.Tag.Lookup("json")
			if found {
				var newname string

				newname, skip, omitempty, nullempty, modifiers = parseTag(tag)

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
			case reflect.Slice,reflect.Map,reflect.Array:
				if reflectedStructFieldValue.Len() <= 0 {
	/////////////////////////////// CONTINUE
					continue
				}
			}
		}

		{
			var valuebytes []byte
			var nulled bool
			if nullempty {
				var empty bool
				var checkedInterface bool

				switch casted := structFieldValueAny.(type) {
				case Emptier:
					empty = casted.IsEmpty()
					checkedInterface = true
				case Nothinger:
					empty = casted.IsNothing()
					checkedInterface = true
				}

				if !checkedInterface {
					var zero reflect.Value = reflect.Zero(reflectedStructFieldType.Type)
					if reflect.DeepEqual(zero.Interface(), structFieldValueAny) {
						empty = true
					}
				}

				if !checkedInterface && !empty {
					switch reflectedStructFieldValue.Kind() {
					case reflect.Slice,reflect.Map,reflect.Array:
						if reflectedStructFieldValue.Len() <= 0 {
							empty = true
						}
					}
				}

				if empty {
					valuebytes = []byte("null")
					nulled = true
				}
			}
			if !nulled {
				var err error
				valuebytes, err = receiver.Marshal(structFieldValueAny)
				if nil != err {
					if _, isErrorEmpty := err.(ErrorEmpty) ; omitempty && isErrorEmpty {
	/////////////////////////////////////// CONTINUE
						continue
					}
					if _, isErrorEmpty := err.(ErrorEmpty) ; nullempty && isErrorEmpty {
						valuebytes = []byte("null")
						goto doneModifiers
					}
					return nil, false, erorr.Errorf("json: problem marshaling %T into JSON: %w", structFieldValueAny, err)
				}

				var omitted bool
				for _, modifierName := range modifiers {
					if "" == modifierName {
						continue
					}

					var fn ModifierFunc
					receiver.modifierFuncs.Let(func(m *map[string]modifierPair){
						if nil == m {
							return
						}
						if nil == *m {
							return
						}

						pair, ok := (*m)[modifierName]
						if ok {
							fn = pair.marshal
						}
					})

					if nil != fn {
						valuebytes, err = fn(valuebytes)
						if nil != err {
							if _, isErrorEmpty := err.(ErrorEmpty) ; omitempty && isErrorEmpty {
								omitted = true
								break
							}
							if _, isErrorEmpty := err.(ErrorEmpty) ; nullempty && isErrorEmpty {
								valuebytes = []byte("null")
								omitted = false
								break
							}
							return nil, false, erorr.Errorf("json: problem marshaling %T into JSON using modifier %q: %w", structFieldValueAny, modifierName, err)
						}
					}
				}
				if omitted {
	/////////////////////////////// CONTINUE
					continue
				}

			}
			doneModifiers:

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

	return p, notempty, nil
}
