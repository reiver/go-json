package json

import (
	"strconv"
)

// MarshalUint64 returns the JSON version of a Go uint64.
func MarshalUint64(value uint64) []byte {
	var result string = strconv.FormatUint(value, 10)
	return []byte(result)
}
