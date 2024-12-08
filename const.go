package json

import (
	"strconv"
	"unsafe"

	"github.com/reiver/go-erorr"
)

type Constantizer interface {
	JSONConst()
	DecodeFromString(string)(any,error)
}

// Const is used to include a constant field in a struct.
//
// For example:
//
//	type AcitivtyLink struct {
//		HRef            string  `json:"href"`
//		MediaType       string  `json:"mediatype"`
//		Rel             string  `json:"rel"`
//		Rev             string  `json:"rev"`
//		Type json.Const[string] `json:"type" json.value:"Link"` // <----------
//	}
//
// Here is another example:
//
//	type Manitoban struct {
//		GivenName               string  `json:"given-name"`
//		AdditionalNames       []string  `json:"additional-names,omitempty"`
//		FamilyName              string  `json:"family-name"`
//		HomeCountry  json.Const[string] `json:"home-country"  json.value:"Canada"`   // <----------
//		HomeProvince json.Const[string] `json:"home-province" json.value:"Manitoba"` // <----------
//		HomeCity                string  `json:"home-city"`
//	}
type Const[T any] struct{}

var _ Constantizer = Const[bool]{}
var _ Constantizer = Const[int]{}
var _ Constantizer = Const[int8]{}
var _ Constantizer = Const[int16]{}
var _ Constantizer = Const[int32]{}
var _ Constantizer = Const[int64]{}
var _ Constantizer = Const[string]{}
var _ Constantizer = Const[uint]{}
var _ Constantizer = Const[uint8]{}
var _ Constantizer = Const[uint16]{}
var _ Constantizer = Const[uint32]{}
var _ Constantizer = Const[uint64]{}

func (Const[T]) JSONConst() {
	// nothing here
}

func (Const[T]) DecodeFromString(str string) (any, error) {
	var t T
	var a any = t

	switch a.(type) {
	case bool:
		return decodeBoolFromString(str)
	case int:
		return decodeIntFromString(str)
	case int8:
		return decodeInt8FromString(str)
	case int16:
		return decodeInt16FromString(str)
	case int32:
		return decodeInt32FromString(str)
	case int64:
		return decodeInt64FromString(str)
	case string:
		return decodeStringFromString(str)
	case uint:
		return decodeUintFromString(str)
	case uint8:
		return decodeUint8FromString(str)
	case uint16:
		return decodeUint16FromString(str)
	case uint32:
		return decodeUint32FromString(str)
	case uint64:
		return decodeUint64FromString(str)
	default:
		return nil, erorr.Errorf("json: cannot decode %t from string %q", a, str)
	}
}

func decodeBoolFromString(str string) (any, error) {
	switch str {
	case "false":
		return false, nil
	case "true":
		return true, nil
	default:
		var nada bool
		return nada, erorr.Errorf("json: problem decoding %T from string %q — not a valid encoding", nada, str)
	}
}

func decodeIntFromString(str string) (any, error) {
	var nada     int
	var bitSize  int= int(unsafe.Sizeof(nada))*8
	var result   int

	i64, err := strconv.ParseInt(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = int(i64)
	return result, nil
}

func decodeInt8FromString(str string) (any, error) {
	var nada     int8
	const bitSize = 8
	var result   int8

	i64, err := strconv.ParseInt(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = int8(i64)
	return result, nil
}

func decodeInt16FromString(str string) (any, error) {
	var nada     int16
	const bitSize = 16
	var result   int16

	i64, err := strconv.ParseInt(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = int16(i64)
	return result, nil
}

func decodeInt32FromString(str string) (any, error) {
	var nada     int32
	const bitSize = 32
	var result   int32

	i64, err := strconv.ParseInt(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = int32(i64)
	return result, nil
}

func decodeInt64FromString(str string) (any, error) {
	var nada     int64
	const bitSize = 64
	var result   int64

	i64, err := strconv.ParseInt(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = i64
	return result, nil
}

func decodeStringFromString(str string) (any, error) {
	return str, nil
}

func decodeUintFromString(str string) (any, error) {
	var nada    uint
	const bitSize int = int(unsafe.Sizeof(nada))*8
	var result  uint

	u64, err := strconv.ParseUint(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = uint(u64)
	return result, nil
}

func decodeUint8FromString(str string) (any, error) {
	var nada    uint8
	const bitSize = 8
	var result  uint8

	u64, err := strconv.ParseUint(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = uint8(u64)
	return result, nil
}

func decodeUint16FromString(str string) (any, error) {
	var nada    uint16
	const bitSize = 16
	var result  uint16

	u64, err := strconv.ParseUint(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = uint16(u64)
	return result, nil
}

func decodeUint32FromString(str string) (any, error) {
	var nada    uint32
	const bitSize = 32
	var result  uint32

	u64, err := strconv.ParseUint(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = uint32(u64)
	return result, nil
}

func decodeUint64FromString(str string) (any, error) {
	var nada    uint64
	const bitSize = 64
	var result  uint64

	u64, err := strconv.ParseUint(str, 10, bitSize)
	if nil != err {
		return nada, erorr.Errorf("json: problem decoding %T from string %q: %w", nada, str, err)
	}

	result = u64
	return result, nil
}
