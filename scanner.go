package json

import (
	"codeberg.org/reiver/go-erorr"
)

// tokenKind represents the type of a JSON token.
type tokenKind int

const (
	tokenObjectBegin tokenKind = iota // {
	tokenObjectEnd                    // }
	tokenArrayBegin                   // [
	tokenArrayEnd                     // ]
	tokenColon                        // :
	tokenComma                        // ,
	tokenString                       // "..."
	tokenNumber                       // -?[0-9]+...
	tokenTrue                         // true
	tokenFalse                        // false
	tokenNull                         // null
	tokenEOF                          // end of input
)

// token represents a single JSON token.
type token struct {
	kind  tokenKind
	value []byte // raw bytes of the token (for string: includes quotes)
	pos   int    // byte offset in the original input
}

// scanner reads JSON tokens from a byte slice.
type scanner struct {
	data       []byte
	pos        int
	length     int
	permissive bool // when true, allow trailing commas, comments, and leading plus-signs
	peeked     *token
}

func newScanner(data []byte, permissive bool) *scanner {
	return &scanner{
		data:       data,
		pos:        0,
		length:     len(data),
		permissive: permissive,
	}
}

// next returns the next token, skipping whitespace (and comments in permissive mode).
func (s *scanner) next() (token, error) {
	if nil != s.peeked {
		tok := *s.peeked
		s.peeked = nil
		return tok, nil
	}

	err := s.skipWhitespace()
	if nil != err {
		return token{}, err
	}

	if s.pos >= s.length {
		return token{kind: tokenEOF, pos: s.pos}, nil
	}

	c := s.data[s.pos]

	switch c {
	case '{':
		tok := token{kind: tokenObjectBegin, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case '}':
		tok := token{kind: tokenObjectEnd, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case '[':
		tok := token{kind: tokenArrayBegin, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case ']':
		tok := token{kind: tokenArrayEnd, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case ':':
		tok := token{kind: tokenColon, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case ',':
		tok := token{kind: tokenComma, value: s.data[s.pos : s.pos+1], pos: s.pos}
		s.pos++
		return tok, nil
	case '"':
		return s.scanString()
	case 't':
		return s.scanLiteral([]byte("true"), tokenTrue)
	case 'f':
		return s.scanLiteral([]byte("false"), tokenFalse)
	case 'n':
		return s.scanLiteral([]byte("null"), tokenNull)
	default:
		if c == '-' || ('0' <= c && c <= '9') || (c == '+' && s.permissive) {
			return s.scanNumber()
		}
		return token{}, erorr.Errorf("json: unexpected character %q at position %d", c, s.pos)
	}
}

// peek returns the next token without consuming it.
func (s *scanner) peek() (token, error) {
	if nil != s.peeked {
		return *s.peeked, nil
	}

	tok, err := s.next()
	if nil != err {
		return token{}, err
	}

	s.peeked = &tok
	return tok, nil
}

// scanString scans a quoted JSON string.
func (s *scanner) scanString() (token, error) {
	start := s.pos
	s.pos++ // skip opening quote

	for s.pos < s.length {
		c := s.data[s.pos]

		if c == '\\' {
			s.pos++ // skip backslash
			if s.pos >= s.length {
				return token{}, erorr.Errorf("json: unterminated string starting at position %d", start)
			}
			// Skip the escaped character.
			// For \uXXXX, skip the 4 hex digits.
			if s.data[s.pos] == 'u' {
				s.pos++ // skip 'u'
				if s.pos+4 > s.length {
					return token{}, erorr.Errorf("json: unterminated \\uXXXX escape at position %d", s.pos-2)
				}
				s.pos += 4
			} else {
				s.pos++
			}
			continue
		}

		if c == '"' {
			s.pos++ // skip closing quote
			return token{kind: tokenString, value: s.data[start:s.pos], pos: start}, nil
		}

		s.pos++
	}

	return token{}, erorr.Errorf("json: unterminated string starting at position %d", start)
}

// scanNumber scans a JSON number.
func (s *scanner) scanNumber() (token, error) {
	start := s.pos

	// Optional leading sign.
	if s.pos < s.length && s.data[s.pos] == '-' {
		s.pos++
	} else if s.pos < s.length && s.data[s.pos] == '+' && s.permissive {
		s.pos++
	}

	// Integer part.
	if s.pos < s.length && s.data[s.pos] == '0' {
		s.pos++
	} else if s.pos < s.length && '1' <= s.data[s.pos] && s.data[s.pos] <= '9' {
		s.pos++
		for s.pos < s.length && '0' <= s.data[s.pos] && s.data[s.pos] <= '9' {
			s.pos++
		}
	} else {
		return token{}, erorr.Errorf("json: invalid number at position %d", start)
	}

	// Fractional part.
	if s.pos < s.length && s.data[s.pos] == '.' {
		s.pos++
		if s.pos >= s.length || s.data[s.pos] < '0' || s.data[s.pos] > '9' {
			return token{}, erorr.Errorf("json: invalid number at position %d — expected digit after decimal point", start)
		}
		for s.pos < s.length && '0' <= s.data[s.pos] && s.data[s.pos] <= '9' {
			s.pos++
		}
	}

	// Exponent part.
	if s.pos < s.length && (s.data[s.pos] == 'e' || s.data[s.pos] == 'E') {
		s.pos++
		if s.pos < s.length && (s.data[s.pos] == '+' || s.data[s.pos] == '-') {
			s.pos++
		}
		if s.pos >= s.length || s.data[s.pos] < '0' || s.data[s.pos] > '9' {
			return token{}, erorr.Errorf("json: invalid number at position %d — expected digit in exponent", start)
		}
		for s.pos < s.length && '0' <= s.data[s.pos] && s.data[s.pos] <= '9' {
			s.pos++
		}
	}

	return token{kind: tokenNumber, value: s.data[start:s.pos], pos: start}, nil
}

// scanLiteral scans a JSON literal (true, false, null).
func (s *scanner) scanLiteral(expected []byte, kind tokenKind) (token, error) {
	start := s.pos

	if s.pos+len(expected) > s.length {
		return token{}, erorr.Errorf("json: unexpected end of input at position %d", start)
	}

	for i, b := range expected {
		if s.data[s.pos+i] != b {
			return token{}, erorr.Errorf("json: invalid literal at position %d", start)
		}
	}

	s.pos += len(expected)
	return token{kind: kind, value: s.data[start:s.pos], pos: start}, nil
}

// skipWhitespace advances past whitespace characters.
// In permissive mode, also skips comments.
func (s *scanner) skipWhitespace() error {
	for s.pos < s.length {
		c := s.data[s.pos]

		switch c {
		case ' ', '\t', '\n', '\r':
			s.pos++
			continue
		}

		if s.permissive && c == '/' {
			if s.pos+1 < s.length {
				next := s.data[s.pos+1]
				if next == '/' {
					// Line comment: skip until newline.
					s.pos += 2
					for s.pos < s.length && s.data[s.pos] != '\n' {
						s.pos++
					}
					continue
				}
				if next == '*' {
					// Block comment: skip until */.
					s.pos += 2
					for {
						if s.pos+1 >= s.length {
							return errUnterminatedComment
						}
						if s.data[s.pos] == '*' && s.data[s.pos+1] == '/' {
							s.pos += 2
							break
						}
						s.pos++
					}
					continue
				}
			}
		}

		break
	}

	return nil
}

// skipValue skips one complete JSON value without allocating.
func (s *scanner) skipValue() error {
	tok, err := s.next()
	if nil != err {
		return err
	}

	switch tok.kind {
	case tokenString, tokenNumber, tokenTrue, tokenFalse, tokenNull:
		return nil
	case tokenObjectBegin:
		return s.skipObject()
	case tokenArrayBegin:
		return s.skipArray()
	default:
		return erorr.Errorf("json: unexpected token at position %d", tok.pos)
	}
}

// skipObject skips the remainder of a JSON object (after the opening brace has been consumed).
func (s *scanner) skipObject() error {
	for {
		tok, err := s.next()
		if nil != err {
			return err
		}
		if tok.kind == tokenObjectEnd {
			return nil
		}

		// Expect a string key.
		if tok.kind != tokenString {
			return erorr.Errorf("json: expected string key at position %d", tok.pos)
		}

		// Expect colon.
		tok, err = s.next()
		if nil != err {
			return err
		}
		if tok.kind != tokenColon {
			return erorr.Errorf("json: expected colon at position %d", tok.pos)
		}

		// Skip the value.
		err = s.skipValue()
		if nil != err {
			return err
		}

		// Expect comma or closing brace.
		tok, err = s.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenComma {
			s.next() // consume comma
		}
	}
}

// skipArray skips the remainder of a JSON array (after the opening bracket has been consumed).
func (s *scanner) skipArray() error {
	// Check for empty array.
	tok, err := s.peek()
	if nil != err {
		return err
	}
	if tok.kind == tokenArrayEnd {
		s.next() // consume ]
		return nil
	}

	for {
		err := s.skipValue()
		if nil != err {
			return err
		}

		tok, err := s.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenArrayEnd {
			s.next() // consume ]
			return nil
		}
		if tok.kind == tokenComma {
			s.next() // consume comma
			continue
		}

		return erorr.Errorf("json: expected comma or ']' at position %d", tok.pos)
	}
}

// scanRawValue reads one complete JSON value and returns its raw bytes.
// This is used for Unmarshaler dispatch and Const[T] validation.
func (s *scanner) scanRawValue() ([]byte, error) {
	start := s.pos

	// Peek to see if we have a peeked token that needs accounting for.
	if nil != s.peeked {
		start = s.peeked.pos
		s.peeked = nil
		s.pos = start
	}

	err := s.skipWhitespace()
	if nil != err {
		return nil, err
	}

	start = s.pos

	err = s.skipValue()
	if nil != err {
		return nil, err
	}

	return s.data[start:s.pos], nil
}
