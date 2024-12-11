package json

import (
	"strconv"
)

// MarshalInt64 returns the JSON version of a Go uint64.
func MarshalInt64(value int64) []byte {
	var result string = strconv.FormatInt(value, 10)
	return []byte(result)
}
