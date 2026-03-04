package json

import (
	"strconv"
	"strings"
)

// MarshalFloat64 returns the JSON version of a Go float64.
func MarshalFloat64(value float64) []byte {
	var result string = strconv.FormatFloat(value, 'f', 18, 64)

	dotIndex := strings.IndexByte(result, '.')
	if 0 <= dotIndex {
		result = strings.TrimRight(result, "0")

		if '.' == result[len(result)-1] {
			result = result[:len(result)-1]
		}
	}

	return []byte(result)
}
