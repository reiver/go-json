package json

import (
	gojson "encoding/json"
)

// MarshalString returns the JSON version of a Go string.
func UnmarshalString(data []byte, dst *string) error {
	return gojson.Unmarshal(data, dst)
}
