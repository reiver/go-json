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
// It ignore the letter-casing of "true" and "false",
// (also) accepts "t" and "f",
// and (also) accepts "1" and "0".
func UnobstructedUnmarshalBool(bytes []byte, dst *bool) error {
	if len(bytes) <= 0 {
		var err error = ErrNotBool
		return erorr.Errorf("json: cannot parse %q as a 'unobstructed' JSON bool: %w", bytes, err)
	}

	var lower []byte = gobytes.ToLower(trim(bytes))
	{
		var byte0 byte = bytes[0]

		switch byte0 {
		case '0', '1', '.':
			lower = NormalizeNumberBytes(lower)
		}
	}

	if 1 == len(lower) {
		switch lower[0] {
		case 'f', '0':
			value := false
			*dst = value
			return nil
		case 't', '1':
			value := true
			*dst = value
			return nil
		}
	}

	{
		err := UnmarshalBool(lower, dst)
		if nil != err {
			return erorr.Errorf("json: cannot parse %q as a 'unobstructed' JSON bool: %w", bytes, err)
		}
	}

	return nil
}
