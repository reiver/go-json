package json

import (
	"encoding"
	gojson "encoding/json"
	"reflect"
)

// Marshal returns the JSON version of 'value'.
//
// omitempty
//
// For Go structs, if a field in the struct includes the struct-tag `omitempty`, then —
// Marshal will NOT include its in the resulting JSON if its Go value is empty.
//
// For example, consider:
//
//	type MyStruct struct {
//		Once   string
//		Twice  string `json:"twice,omitempty"` // <---------
//		Thrice string `json:"thrice"`
//		Fource string `json:",omitempty"`      // <---------
//	}
//
// Note that field `Twice` and field `Fource` both have `omitempty` in their struct-tags.
// So, if their values are empty, then the resulting JSON will omit them.
//
// For example, this:
//
//	var value MyStruct
//
// Would (conceptually) result in:
//
//	{
//		"Once":   "",
//		"thrice": ""
//	}
//
// And, for example, this:
//
//	var value = MyStruct{
//		Once:   "",
//		Twice:  "",
//		Thrice: "",
//		Fource: ""
//	}
//
// Would also (conceptually) result in:
//
//	{
//		"Once":   "",
//		"thrice": ""
//	}
//
// And also, for example, this:
//
//	var value = MyStruct{
//		Once:   "first",
//		Twice:  "second",
//		Thrice: "third",
//		Fource: "fourth"
//	}
//
// Would (conceptually) result in:
//
//	{
//		"Once":   "first",
//		"twice":  "second",
//		"thrice": "third",
//		"Fource": "forth"
//	}
//
// Custom types can also make use of [Emptier] or [Nothinger] to specify when they are empty.
// For example:
//
//	type MyStruct struct {
//		// ...
//	}
//	
//	func (receiver MyStruct) IsEmpty() bool {
//		// ...
//	}
//
// Marshal will call IsEmpty, if a custom type has it, to check whether the custom type is `empty` or not, for the purposes of `omitempty`.
func (receiver *Usher) Marshal(value any) ([]byte, error) {
	if nil == value {
		return []byte{'n','u','l','l'}, nil
	}

	switch casted := value.(type) {
	case Marshaler:
		return casted.MarshalJSON()
	case encoding.TextMarshaler:
		return MarshalTextMarshaler(casted)
	case string:
		return MarshalString(casted), nil
	case bool:
		return MarshalBool(casted), nil
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
	case uint:
		return MarshalUint(casted), nil
	case uint8:
		return MarshalUint8(casted), nil
	case uint16:
		return MarshalUint16(casted), nil
	case uint32:
		return MarshalUint32(casted), nil
	case uint64:
		return MarshalUint64(casted), nil

	default:
		reflectedType := reflect.TypeOf(value)
		if nil == reflectedType {
			return nil, errBadReflection
		}

		switch reflectedType.Kind() {
		case reflect.Struct:
			return receiver.marshalStruct(value)
//@TODO: array
		case reflect.Slice:
			return marshalSlice(value, receiver.Marshal)
		case reflect.Map:
			return receiver.marshalMap(value)
		default:
			return gojson.Marshal(value)
		}
	}
}
