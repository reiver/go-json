package json

import (
	"unicode/utf8"

	"codeberg.org/reiver/go-erorr"
)

// unquoteString takes a JSON string token (including the surrounding quotes)
// and returns the decoded Go string.
func unquoteString(data []byte) (string, error) {
	if len(data) < 2 {
		return "", errNotStringMissingBeginQuotationMark
	}
	if data[0] != '"' {
		return "", errNotStringMissingBeginQuotationMark
	}
	if data[len(data)-1] != '"' {
		return "", errNotStringMissingEndQuotationMark
	}

	data = data[1 : len(data)-1]

	// Fast path: if no backslash, return directly.
	hasEscape := false
	for _, b := range data {
		if b == '\\' {
			hasEscape = true
			break
		}
	}
	if !hasEscape {
		return string(data), nil
	}

	// Slow path: process escape sequences.
	buf := make([]byte, 0, len(data))
	i := 0
	for i < len(data) {
		if data[i] != '\\' {
			buf = append(buf, data[i])
			i++
			continue
		}

		i++ // skip the backslash
		if i >= len(data) {
			return "", errInvalidEscape
		}

		switch data[i] {
		case '"':
			buf = append(buf, '"')
			i++
		case '\\':
			buf = append(buf, '\\')
			i++
		case '/':
			buf = append(buf, '/')
			i++
		case 'b':
			buf = append(buf, '\b')
			i++
		case 'f':
			buf = append(buf, '\f')
			i++
		case 'n':
			buf = append(buf, '\n')
			i++
		case 'r':
			buf = append(buf, '\r')
			i++
		case 't':
			buf = append(buf, '\t')
			i++
		case 'u':
			i++ // skip 'u'
			if i+4 > len(data) {
				return "", erorr.Errorf("json: invalid \\uXXXX escape at position %d", i-2)
			}
			r1, err := parseHex4(data[i : i+4])
			if nil != err {
				return "", err
			}
			i += 4

			// Check for surrogate pair.
			if 0xD800 <= r1 && r1 <= 0xDBFF {
				if i+6 > len(data) || data[i] != '\\' || data[i+1] != 'u' {
					return "", erorr.Errorf("json: invalid surrogate pair at position %d", i)
				}
				i += 2 // skip \u
				r2, err := parseHex4(data[i : i+4])
				if nil != err {
					return "", err
				}
				i += 4

				if r2 < 0xDC00 || r2 > 0xDFFF {
					return "", erorr.Errorf("json: invalid low surrogate at position %d", i-4)
				}

				combined := rune(0x10000) + rune(r1-0xD800)*0x400 + rune(r2-0xDC00)
				var tmp [4]byte
				n := utf8.EncodeRune(tmp[:], rune(combined))
				buf = append(buf, tmp[:n]...)
			} else if 0xDC00 <= r1 && r1 <= 0xDFFF {
				return "", erorr.Errorf("json: invalid lone low surrogate at position %d", i-4)
			} else {
				var tmp [4]byte
				n := utf8.EncodeRune(tmp[:], rune(r1))
				buf = append(buf, tmp[:n]...)
			}
		default:
			return "", erorr.Errorf("json: invalid escape character %q at position %d", data[i], i)
		}
	}

	return string(buf), nil
}

// parseHex4 parses a 4-character hexadecimal sequence into a uint16.
func parseHex4(data []byte) (uint16, error) {
	var result uint16

	for i := 0; i < 4; i++ {
		var digit uint16

		c := data[i]
		switch {
		case '0' <= c && c <= '9':
			digit = uint16(c - '0')
		case 'a' <= c && c <= 'f':
			digit = uint16(c-'a') + 10
		case 'A' <= c && c <= 'F':
			digit = uint16(c-'A') + 10
		default:
			return 0, erorr.Errorf("json: invalid hex character %q in \\uXXXX escape", c)
		}

		result = result*16 + digit
	}

	return result, nil
}
