package json

// NormalizeNumberString normalizes a JSON number.
// It expands scientific notation into plain decimal form (e.g., "3.7e-5" becomes "0.000037"),
// strips leading '+', unnecessary leading zeros, trailing zeros after a decimal point, and a trailing decimal point.
// A negative sign is preserved when appropriate.
//
// See also: [Number].
func NormalizeNumberString(str string) string {
	if len(str) == 0 {
		return "0"
	}

	// Track and strip sign.
	negative := false
	s := str
	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 {
		return "0"
	}

	// Expand scientific notation (e.g., "3.7e-5" → "0.000037").
	ePos := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'e' || s[i] == 'E' {
			ePos = i
			break
		}
	}
	if ePos >= 0 {
		mantissa := s[:ePos]
		expStr := s[ePos+1:]

		// Parse exponent.
		expNeg := false
		if len(expStr) > 0 && expStr[0] == '-' {
			expNeg = true
			expStr = expStr[1:]
		} else if len(expStr) > 0 && expStr[0] == '+' {
			expStr = expStr[1:]
		}
		exp := 0
		for i := 0; i < len(expStr); i++ {
			exp = exp*10 + int(expStr[i]-'0')
		}
		if expNeg {
			exp = -exp
		}

		// Extract digits from mantissa and find the decimal position.
		var digits []byte
		decPos := -1
		for i := 0; i < len(mantissa); i++ {
			if mantissa[i] == '.' {
				decPos = len(digits)
			} else {
				digits = append(digits, mantissa[i])
			}
		}
		if decPos < 0 {
			decPos = len(digits)
		}

		// Shift decimal position by exponent.
		newDecPos := decPos + exp

		// Rebuild the plain decimal string.
		if newDecPos <= 0 {
			// All digits are fractional, prepend "0." and leading zeros.
			var buf []byte
			buf = append(buf, '0', '.')
			for i := 0; i < -newDecPos; i++ {
				buf = append(buf, '0')
			}
			buf = append(buf, digits...)
			s = string(buf)
		} else if newDecPos >= len(digits) {
			// All digits are integer, append trailing zeros.
			var buf []byte
			buf = append(buf, digits...)
			for i := 0; i < newDecPos-len(digits); i++ {
				buf = append(buf, '0')
			}
			s = string(buf)
		} else {
			// Decimal point falls within the digits.
			var buf []byte
			buf = append(buf, digits[:newDecPos]...)
			buf = append(buf, '.')
			buf = append(buf, digits[newDecPos:]...)
			s = string(buf)
		}
	}

	// Find decimal point position (if any).
	dotPos := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			dotPos = i
			break
		}
	}

	var intPart, fracPart string
	if dotPos >= 0 {
		intPart = s[:dotPos]
		fracPart = s[dotPos+1:]
	} else {
		intPart = s
	}

	// Strip leading zeros from integer part, keeping at least one digit.
	i := 0
	for i < len(intPart)-1 && intPart[i] == '0' {
		i++
	}
	intPart = intPart[i:]

	if len(intPart) == 0 {
		intPart = "0"
	}

	// Strip trailing zeros from fractional part.
	if len(fracPart) > 0 {
		j := len(fracPart)
		for j > 0 && fracPart[j-1] == '0' {
			j--
		}
		fracPart = fracPart[:j]
	}

	// Build result.
	var result string
	if len(fracPart) > 0 {
		result = intPart + "." + fracPart
	} else {
		result = intPart
	}

	// Check if result is just "0".
	if result == "0" {
		return "0"
	}

	if negative {
		return "-" + result
	}
	return result
}
