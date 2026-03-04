package json

import (
	"strconv"
	"strings"
)

// MarshalFloat32 returns the JSON version of a Go float32.
func MarshalFloat32(value float32) []byte {
	var result string = strconv.FormatFloat(float64(value), 'f', 18, 32)

	dotIndex := strings.IndexByte(result, '.')
	if 0 <= dotIndex {
		result = strings.TrimRight(result, "0")

		if '.' == result[len(result)-1] {
			result = result[:len(result)-1]
		}
	}

	return []byte(result)
}
