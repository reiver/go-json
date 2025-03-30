package json

import (
	gojson "encoding/json"
)

func Unmarshal(data []byte, dst any) error {
	return gojson.Unmarshal(data, dst)
}
