package json

import (
	gobytes "bytes"

	"codeberg.org/reiver/go-erorr"
)

var (
	falseBytes []byte = []byte("false")
	trueBytes  []byte = []byte("true")
)

// UnmarshalBool JSON-unmarshals a JSON bool into a Go bool.
func UnmarshalBool(bytes []byte, dst *bool) error {
	bytes = trim(bytes)

	var value bool
	switch {
	case gobytes.Equal(bytes, falseBytes):
		value = false
	case gobytes.Equal(bytes, trueBytes):
		value = true
	default:
		var err error = ErrNotBool
		return erorr.Errorf("json: cannot parse %q as a JSON bool: %w", bytes, err)
	}

	*dst = value
	return nil
}

// UnobstructedUnmarshal JSON-unmarshals a JSON bool into a Go bool.
// It ignore the letter-casing of "true" and "false".
func UnobstructedUnmarshalBool(bytes []byte, dst *bool) error {
	var lower []byte = gobytes.ToLower(bytes)
	err := UnmarshalBool(lower, dst)
	if nil != err {
		return erorr.Errorf("json: cannot parse %q as a 'unobstructed' JSON bool: %w", bytes, err)
	}

	return nil
}
