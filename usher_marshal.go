package json

import (
	"encoding"
	gojson "encoding/json"
	"reflect"
)

// Marshal returns the JSON version of 'value'.
func (receiver *Usher) Marshal(value any) ([]byte, error) {
	if nil == value {
		return []byte{'n','u','l','l'}, nil
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
	case string:
		return MarshalString(casted), nil
	case int:
		return MarshalInt(casted), nil
	case int8:
		return MarshalInt8(casted), nil
	case int16:
		return MarshalInt16(casted), nil
	case int32:
		return MarshalInt32(casted), nil
	case int64:
		return MarshalInt64(casted), nil

	default:
		reflectedType := reflect.TypeOf(value)
		if nil == reflectedType {
			return nil, errBadReflection
		}

		switch reflectedType.Kind() {
//@TODO: array
		case reflect.Struct:
			return receiver.marshalStruct(value)
		case reflect.Slice:
			return receiver.marshalSlice(value)
		case reflect.Map:
			return receiver.marshalMap(value)
		default:
			return gojson.Marshal(value)
		}
	}
}
