package json

import (
	gobytes "bytes"
)

var (
	falseBytes []byte = []byte("false")
	trueBytes  []byte = []byte("true")
)

// UnmarshalBool JSON-unmarshals a JSON bool into a Go bool.
func UnmarshalBool(bytes []byte, dst *bool) error {
	const spacing string = " \t\n\r"
	bytes = gobytes.Trim(bytes, spacing)

	var value bool
	switch {
	case gobytes.Equal(bytes, falseBytes):
		value = false
	case gobytes.Equal(bytes, trueBytes):
		value = true
	default:
		return ErrNotBool
	}

	*dst = value
	return nil
}
