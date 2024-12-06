package json

import (
	"reflect"

	"github.com/reiver/go-erorr"
)

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
