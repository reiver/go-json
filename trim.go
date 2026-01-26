package json

import (
	gobytes "bytes"
)

func trim(bytes []byte) []byte {
	const spacing string = " \t\n\r"
	return gobytes.Trim(bytes, spacing)
}
