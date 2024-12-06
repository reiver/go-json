package json

import (
	"encoding"
	"fmt"
	"reflect"
	"sort"

	"github.com/reiver/go-erorr"
)

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
