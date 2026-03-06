package json

import (
	"reflect"
	"strconv"
)

type unsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func tryParseJSONUint[T unsignedInt](data []byte) (T, error) {
	var zero T
	bitSize := int(reflect.TypeOf(zero).Size()) * 8
	n, err := strconv.ParseUint(string(data), 10, bitSize)
	if nil != err {
		var nada T
		return nada, err
	}
	return T(n), nil
}

func parseJSONUint[T unsignedInt](data []byte) (any, error) {
	result, err := tryParseJSONUint[T](data)
	if nil != err {
		return nil, err
	}

	return result, nil
}

func UnmarshalUint(data []byte, dst *uint) error {
	datum, err := tryParseJSONUint[uint](data)
	if nil != err {
		return err
	}

	if nil == dst {
		var space uint
		dst = &space
	}

	*dst = datum
	return nil
}

func UnmarshalUint8(data []byte, dst *uint8) error {
	datum, err := tryParseJSONUint[uint8](data)
	if nil != err {
		return err
	}

	if nil == dst {
		var space uint8
		dst = &space
	}

	*dst = datum
	return nil
}

func UnmarshalUint16(data []byte, dst *uint16) error {
	datum, err := tryParseJSONUint[uint16](data)
	if nil != err {
		return err
	}

	if nil == dst {
		var space uint16
		dst = &space
	}

	*dst = datum
	return nil
}

func UnmarshalUint32(data []byte, dst *uint32) error {
	datum, err := tryParseJSONUint[uint32](data)
	if nil != err {
		return err
	}

	if nil == dst {
		var space uint32
		dst = &space
	}

	*dst = datum
	return nil
}

func UnmarshalUint64(data []byte, dst *uint64) error {
	datum, err := tryParseJSONUint[uint64](data)
	if nil != err {
		return err
	}

	if nil == dst {
		var space uint64
		dst = &space
	}

	*dst = datum
	return nil
}
