package json

import (
	"reflect"
	"strconv"
)

type signedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func tryParseJSONInt[T signedInt](data []byte) (T, error) {
	var zero T
	bitSize := int(reflect.TypeOf(zero).Size()) * 8
	n, err := strconv.ParseInt(string(data), 10, bitSize)
	if nil != err {
		var nada T
		return nada, err
	}
	return T(n), nil
}

func parseJSONInt[T signedInt](data []byte) (any, error) {
	result, err := tryParseJSONInt[T](data)
	if nil != err {
		return nil, err
	}

	return result, nil
}

func UnmarshalInt(data []byte, dst *int) error {
        datum, err := tryParseJSONInt[int](data)
        if nil != err {
                return err
        }

        if nil == dst {
                var space int
                dst = &space
        }

        *dst = datum
        return nil
}

func UnmarshalInt8(data []byte, dst *int8) error {
        datum, err := tryParseJSONInt[int8](data)
        if nil != err {
                return err
        }

        if nil == dst {
                var space int8
                dst = &space
        }

        *dst = datum
        return nil
}

func UnmarshalInt16(data []byte, dst *int16) error {
        datum, err := tryParseJSONInt[int16](data)
        if nil != err {
                return err
        }

        if nil == dst {
                var space int16
                dst = &space
        }

        *dst = datum
        return nil
}

func UnmarshalInt32(data []byte, dst *int32) error {
        datum, err := tryParseJSONInt[int32](data)
        if nil != err {
                return err
        }

        if nil == dst {
                var space int32
                dst = &space
        }

        *dst = datum
        return nil
}

func UnmarshalInt64(data []byte, dst *int64) error {
        datum, err := tryParseJSONInt[int64](data)
        if nil != err {
                return err
        }

        if nil == dst {
                var space int64
                dst = &space
        }

        *dst = datum
        return nil
}
