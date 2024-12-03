package json

import (
	gojson "encoding/json"
)

// MarshalString returns the JSON version of a Go string.
func MarshalString(value string) []byte {
	result, err := gojson.Marshal(value)
	if nil != err {
		panic(err)
	}

	return result
}
