package json

import (
	"encoding"
)

// MarshalTextMarshaler returns the JSON version of the successful result from encoding.TextMarshaler..
func MarshalTextMarshaler(textMarshaler encoding.TextMarshaler) ([]byte, error) {
	bytes, err := textMarshaler.MarshalText()
	if nil != err {
		return nil, err
	}

	var str string = string(bytes)
	return MarshalString(str), nil
}
