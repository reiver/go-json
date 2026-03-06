package json

import (
	"strconv"

	"codeberg.org/reiver/go-erorr"
)

// Number represents a JSON number.
// It stores it as a string to avoid precision loss that comes from using float64.
//
// See also: [NormalizeNumber].
type Number struct {
	value string
}

// ParseNumberString returns a [Number] with the value of a JSON number.
// The second return value is false if jsonNumber is not a valid JSON number.
func ParseNumberString(jsonNumber string) (Number, bool) {
	if !isJSONNumber(jsonNumber) {
		return Number{}, false
	}

	var num Number
	num.set(jsonNumber)
	return num, true
}

func MustParseNumberString(jsonNumber string) Number {
	num, ok := ParseNumberString(jsonNumber)
	if !ok {
		var err error = erorr.Errorf("json: failed to parse string %q as JSON number: %w", jsonNumber, ErrNotJSONNumber)
		panic(err)
	}

	return num
}

// isJSONNumber reports whether s is a valid JSON number literal.
//
// JSON number grammar:
//
//	number = [ '-' ] int [ frac ] [ exp ]
//	int    = '0' | ( '1'..'9' { '0'..'9' } )
//	frac   = '.' '0'..'9' { '0'..'9' }
//	exp    = ( 'e' | 'E' ) [ '+' | '-' ] '0'..'9' { '0'..'9' }
func isJSONNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0

	// Optional leading minus.
	if s[i] == '-' {
		i++
		if i >= len(s) {
			return false
		}
	}

	// Integer part.
	if s[i] == '0' {
		i++
	} else if '1' <= s[i] && s[i] <= '9' {
		i++
		for i < len(s) && '0' <= s[i] && s[i] <= '9' {
			i++
		}
	} else {
		return false
	}

	// Optional fractional part.
	if i < len(s) && s[i] == '.' {
		i++
		if i >= len(s) || s[i] < '0' || s[i] > '9' {
			return false
		}
		for i < len(s) && '0' <= s[i] && s[i] <= '9' {
			i++
		}
	}

	// Optional exponent part.
	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		i++
		if i >= len(s) {
			return false
		}
		if s[i] == '+' || s[i] == '-' {
			i++
			if i >= len(s) {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		for i < len(s) && '0' <= s[i] && s[i] <= '9' {
			i++
		}
	}

	return i == len(s)
}

// Zero returns a [Number] with a value of zero (0).
func Zero() Number {
	var zero Number
	return zero
}

// get returns a string representation of the number.
// If the value is empty, it returns "0".
func (receiver Number) get() string {
	if "" == receiver.value {
		return "0"
	}
	return receiver.value
}

// set normalizes and stores the given numeric string.
// If the normalized value is "0", it stores "" to make it so an uninitialized [Number] has the same value as one set to zero.
func (receiver *Number) set(value string) {
	normalized := NormalizeNumber(value)
	if "0" == normalized {
		receiver.value = ""
		return
	}
	receiver.value = normalized
}

// String returns the string representation of the number.
//
// String makes [Number] fit the [fmt.Stringer] interface.
func (receiver Number) String() string {
	return receiver.get()
}

// MarshalJSON returns the JSON encoding of the number.
//
// String makes [Number] fit the [Marshaler] interface.
func (receiver Number) MarshalJSON() ([]byte, error) {
	return []byte(receiver.get()), nil
}

// Int64 returns the number as an int64.
// The second return value is false if the number cannot be represented as an int64.
func (receiver Number) Int64() (int64, bool) {
	n, err := strconv.ParseInt(receiver.get(), 10, 64)
	if nil != err {
		return 0, false
	}
	return n, true
}

// Uint64 returns the number as a uint64.
// The second return value is false if the number cannot be represented as a uint64.
func (receiver Number) Uint64() (uint64, bool) {
	n, err := strconv.ParseUint(receiver.get(), 10, 64)
	if nil != err {
		return 0, false
	}
	return n, true
}

// Float64 returns the number as a float64.
// The second return value is false if the number cannot be represented as a float64.
func (receiver Number) Float64() (float64, bool) {
	f, err := strconv.ParseFloat(receiver.get(), 64)
	if nil != err {
		return 0, false
	}
	return f, true
}

// UnmarshalJSON sets the number from a JSON number literal.
// Returns an error if data is not a valid JSON number.
//
// UnmarshalJSON makes [Number] fit the [Unmarshaler] interface.
func (receiver *Number) UnmarshalJSON(data []byte) error {
	var str string = string(data)

	if !isJSONNumber(str) {
		return erorr.Errorf("json: cannot unmarshal %q into Number: %w", str, ErrNotJSONNumber)
	}
	receiver.set(str)
	return nil
}
