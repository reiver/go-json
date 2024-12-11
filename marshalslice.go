package json

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

// MarshalString returns the JSON version of a Go string.
func MarshalSlice[T any](value []T) ([]byte, error) {
	return marshalSlice(value, Marshal)
}

func marshalSlice(value any, marshalFunc func(any)([]byte, error)) ([]byte, error) {
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

			bytes, err := marshalFunc(element)
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
