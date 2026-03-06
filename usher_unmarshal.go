package json

import (
	"reflect"
	"strconv"

	"codeberg.org/reiver/go-erorr"
)

// Unmarshal deserializes JSON data into the value pointed to by dst.
// Unknown JSON fields are silently ignored.
func (receiver *Usher) Unmarshal(data []byte, dst any) error {
	return receiver.unmarshalWithMode(data, dst, unmarshalModeStandard)
}

// UnobstructedUnmarshal deserializes JSON data permissively.
// Allows trailing commas, comments, and leading plus-signs.
// Unknown JSON fields are silently ignored.
func (receiver *Usher) UnobstructedUnmarshal(data []byte, dst any) error {
	return receiver.unmarshalWithMode(data, dst, unmarshalModeUnobstructed)
}

// ObstructedUnmarshal deserializes JSON data strictly.
// Returns an error if unknown JSON fields are encountered.
func (receiver *Usher) ObstructedUnmarshal(data []byte, dst any) error {
	return receiver.unmarshalWithMode(data, dst, unmarshalModeObstructed)
}

func (receiver *Usher) unmarshalWithMode(data []byte, dst any, mode unmarshalMode) error {
	if nil == dst {
		return errNilDestination
	}

	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Pointer {
		return errNotPointer
	}
	if rv.IsNil() {
		return errNilDestination
	}

	permissive := mode == unmarshalModeUnobstructed
	sc := newScanner(data, permissive)

	var path jsonPath
	var errs []error

	err := receiver.unmarshalValue(sc, rv.Elem(), mode, &path, &errs)
	if nil != err {
		return err
	}

	// Check for trailing data.
	tok, err := sc.peek()
	if nil != err {
		return err
	}
	if tok.kind != tokenEOF {
		return erorr.Errorf("json: trailing data at position %d", tok.pos)
	}

	if len(errs) > 0 {
		return UnmarshalErrors{Errors: errs}
	}

	return nil
}

// unmarshalValue is the recursive dispatch function.
func (receiver *Usher) unmarshalValue(
	sc *scanner,
	dst reflect.Value,
	mode unmarshalMode,
	path *jsonPath,
	errs *[]error,
) error {

	// Handle null before pointer dereferencing — set pointer to nil.
	if dst.Kind() == reflect.Pointer {
		tok, err := sc.peek()
		if nil != err {
			return erorr.Errorf("json: problem reading token at %s: %w", path.String(), err)
		}
		if tok.kind == tokenNull {
			sc.next() // consume null
			dst.Set(reflect.Zero(dst.Type()))
			return nil
		}
	}

	// Handle pointers: allocate if nil, recurse into Elem().
	for dst.Kind() == reflect.Pointer {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		dst = dst.Elem()
	}

	// Check if dst implements Unmarshaler (via pointer).
	if dst.CanAddr() {
		addr := dst.Addr()
		if u, ok := addr.Interface().(Unmarshaler); ok {
			raw, err := sc.scanRawValue()
			if nil != err {
				return erorr.Errorf("json: problem reading value at %s: %w", path.String(), err)
			}
			err = u.UnmarshalJSON(raw)
			if nil != err {
				return erorr.Errorf("json: problem unmarshaling %s into %T: %w", path.String(), u, err)
			}
			return nil
		}
	}

	// Peek at the next token.
	tok, err := sc.peek()
	if nil != err {
		return erorr.Errorf("json: problem reading token at %s: %w", path.String(), err)
	}

	// Handle interface{}/any: unmarshal into natural Go types.
	if dst.Kind() == reflect.Interface {
		val, err := receiver.unmarshalAny(sc, mode, path, errs)
		if nil != err {
			return err
		}
		if nil != val {
			dst.Set(reflect.ValueOf(val))
		} else {
			dst.Set(reflect.Zero(dst.Type()))
		}
		return nil
	}

	switch tok.kind {
	case tokenNull:
		sc.next() // consume null
		dst.Set(reflect.Zero(dst.Type()))
		return nil

	case tokenString:
		sc.next() // consume the string token
		str, err := unquoteString(tok.value)
		if nil != err {
			return erorr.Errorf("json: problem decoding string at %s: %w", path.String(), err)
		}
		switch dst.Kind() {
		case reflect.String:
			dst.SetString(str)
		default:
			return erorr.Errorf("json: cannot unmarshal string into %s at %s", dst.Type(), path.String())
		}
		return nil

	case tokenNumber:
		sc.next() // consume the number token
		return unmarshalNumber(tok, dst, path)

	case tokenTrue:
		sc.next()
		if dst.Kind() != reflect.Bool {
			return erorr.Errorf("json: cannot unmarshal bool into %s at %s", dst.Type(), path.String())
		}
		dst.SetBool(true)
		return nil

	case tokenFalse:
		sc.next()
		if dst.Kind() != reflect.Bool {
			return erorr.Errorf("json: cannot unmarshal bool into %s at %s", dst.Type(), path.String())
		}
		dst.SetBool(false)
		return nil

	case tokenObjectBegin:
		switch dst.Kind() {
		case reflect.Struct:
			return receiver.unmarshalStruct(sc, dst, mode, path, errs)
		case reflect.Map:
			return receiver.unmarshalMap(sc, dst, mode, path, errs)
		default:
			return erorr.Errorf("json: cannot unmarshal object into %s at %s", dst.Type(), path.String())
		}

	case tokenArrayBegin:
		switch dst.Kind() {
		case reflect.Slice:
			return receiver.unmarshalSlice(sc, dst, mode, path, errs)
		case reflect.Array:
			return receiver.unmarshalArray(sc, dst, mode, path, errs)
		default:
			return erorr.Errorf("json: cannot unmarshal array into %s at %s", dst.Type(), path.String())
		}

	default:
		return erorr.Errorf("json: unexpected token at position %d", tok.pos)
	}
}

// unmarshalNumber sets a numeric value on dst from a JSON number token.
func unmarshalNumber(tok token, dst reflect.Value, path *jsonPath) error {
	numStr := string(tok.value)

	switch dst.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		normalized := NormalizeNumber(numStr)
		n, err := strconv.ParseInt(normalized, 10, dst.Type().Bits())
		if nil != err {
			return erorr.Errorf("json: cannot unmarshal %q into %s at %s: %w", numStr, dst.Type(), path.String(), err)
		}
		dst.SetInt(n)
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		normalized := NormalizeNumber(numStr)
		n, err := strconv.ParseUint(normalized, 10, dst.Type().Bits())
		if nil != err {
			return erorr.Errorf("json: cannot unmarshal %q into %s at %s: %w", numStr, dst.Type(), path.String(), err)
		}
		dst.SetUint(n)
		return nil

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(numStr, dst.Type().Bits())
		if nil != err {
			return erorr.Errorf("json: cannot unmarshal %q into %s at %s: %w", numStr, dst.Type(), path.String(), err)
		}
		dst.SetFloat(f)
		return nil

	default:
		return erorr.Errorf("json: cannot unmarshal number into %s at %s", dst.Type(), path.String())
	}
}

// unmarshalAny unmarshals a JSON value into a natural Go type (for interface{}/any destinations).
func (receiver *Usher) unmarshalAny(
	sc *scanner,
	mode unmarshalMode,
	path *jsonPath,
	errs *[]error,
) (any, error) {
	tok, err := sc.peek()
	if nil != err {
		return nil, err
	}

	switch tok.kind {
	case tokenNull:
		sc.next()
		return nil, nil

	case tokenTrue:
		sc.next()
		return true, nil

	case tokenFalse:
		sc.next()
		return false, nil

	case tokenString:
		sc.next()
		str, err := unquoteString(tok.value)
		if nil != err {
			return nil, err
		}
		return str, nil

	case tokenNumber:
		sc.next()
		var num Number
		num.set(string(tok.value))
		return num, nil

	case tokenObjectBegin:
		sc.next() // consume {
		m := make(map[string]any)

		for {
			tok, err := sc.peek()
			if nil != err {
				return nil, err
			}
			if tok.kind == tokenObjectEnd {
				sc.next() // consume }
				return m, nil
			}

			// Read key.
			keyTok, err := sc.next()
			if nil != err {
				return nil, err
			}
			if keyTok.kind != tokenString {
				return nil, erorr.Errorf("json: expected string key at position %d", keyTok.pos)
			}
			key, err := unquoteString(keyTok.value)
			if nil != err {
				return nil, err
			}

			// Read colon.
			colonTok, err := sc.next()
			if nil != err {
				return nil, err
			}
			if colonTok.kind != tokenColon {
				return nil, erorr.Errorf("json: expected colon at position %d", colonTok.pos)
			}

			// Read value.
			path.pushKey(key)
			val, err := receiver.unmarshalAny(sc, mode, path, errs)
			if nil != err {
				return nil, err
			}
			path.pop()

			m[key] = val

			// Comma or closing brace.
			tok, err = sc.peek()
			if nil != err {
				return nil, err
			}
			if tok.kind == tokenComma {
				sc.next()

				// Handle trailing comma in permissive mode.
				if mode == unmarshalModeUnobstructed {
					tok, err = sc.peek()
					if nil != err {
						return nil, err
					}
					if tok.kind == tokenObjectEnd {
						sc.next()
						return m, nil
					}
				}
			}
		}

	case tokenArrayBegin:
		sc.next() // consume [
		var arr []any

		for {
			tok, err := sc.peek()
			if nil != err {
				return nil, err
			}
			if tok.kind == tokenArrayEnd {
				sc.next() // consume ]
				return arr, nil
			}

			path.pushIndex(len(arr))
			val, err := receiver.unmarshalAny(sc, mode, path, errs)
			if nil != err {
				return nil, err
			}
			path.pop()

			arr = append(arr, val)

			// Comma or closing bracket.
			tok, err = sc.peek()
			if nil != err {
				return nil, err
			}
			if tok.kind == tokenComma {
				sc.next()

				// Handle trailing comma in permissive mode.
				if mode == unmarshalModeUnobstructed {
					tok, err = sc.peek()
					if nil != err {
						return nil, err
					}
					if tok.kind == tokenArrayEnd {
						sc.next()
						return arr, nil
					}
				}
			}
		}

	default:
		return nil, erorr.Errorf("json: unexpected token at position %d", tok.pos)
	}
}
