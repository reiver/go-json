package json

import (
	gojson "encoding/json"
)

// UnmarshalString JSON-unmarshals a JSON string into a Go string.
func UnmarshalString(data []byte, dst *string) error {
	return gojson.Unmarshal(data, dst)
}
