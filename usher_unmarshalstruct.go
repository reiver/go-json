package json

import (
	"reflect"
	"strconv"

	"codeberg.org/reiver/go-erorr"
)

// unmarshalStruct unmarshals a JSON object into a Go struct.
// The opening '{' has NOT yet been consumed.
func (receiver *Usher) unmarshalStruct(
	sc *scanner,
	dst reflect.Value,
	mode unmarshalMode,
	path *jsonPath,
	errs *[]error,
) error {
	// Consume the opening brace.
	tok, err := sc.next()
	if nil != err {
		return err
	}
	if tok.kind != tokenObjectBegin {
		return erorr.Errorf("json: expected '{' at position %d", tok.pos)
	}

	info := getStructInfo(dst.Type())

	for {
		// Peek to check for closing brace or next key.
		tok, err := sc.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenObjectEnd {
			sc.next() // consume }
			return nil
		}

		// Read the key.
		keyTok, err := sc.next()
		if nil != err {
			return err
		}
		if keyTok.kind != tokenString {
			return erorr.Errorf("json: expected string key at position %d", keyTok.pos)
		}
		keyName, err := unquoteString(keyTok.value)
		if nil != err {
			return err
		}

		// Read the colon.
		colonTok, err := sc.next()
		if nil != err {
			return err
		}
		if colonTok.kind != tokenColon {
			return erorr.Errorf("json: expected colon at position %d", colonTok.pos)
		}

		path.pushKey(keyName)

		// Look up the field.
		fieldIdx, found := info.nameIndex[keyName]

		if !found {
			// Unknown field.
			if mode == unmarshalModeObstructed {
				*errs = append(*errs, UnknownFieldError{
					Path: path.String(),
					Key:  keyName,
				})
			}
			err = sc.skipValue()
			if nil != err {
				return err
			}
			path.pop()
			goto readSeparator
		}

		{
			sf := info.fields[fieldIdx]

			// Handle OmitAlways fields: skip the value.
			if sf.isOmitAlways {
				err = sc.skipValue()
				if nil != err {
					return err
				}
				path.pop()
				goto readSeparator
			}

			// Handle Const[T] fields: validate value matches expected constant.
			if sf.isConst {
				err = receiver.validateConst(sc, dst, sf, path, errs)
				if nil != err {
					return err
				}
				path.pop()
				goto readSeparator
			}

			// Regular field: recursively unmarshal.
			fieldValue := dst.Field(sf.index)
			if len(sf.modifiers) > 0 {
				// Apply reverse modifiers: read the raw value, transform it, then unmarshal.
				raw, rawErr := sc.scanRawValue()
				if nil != rawErr {
					return rawErr
				}

				for i := len(sf.modifiers) - 1; i >= 0; i-- {
					modifierName := sf.modifiers[i]
					if "" == modifierName {
						continue
					}

					var fn ModifierFunc
					var found bool
					receiver.modifierFuncs.Let(func(m *map[string]modifierPair){
						if nil == m {
							return
						}
						if nil == *m {
							return
						}

						pair, ok := (*m)[modifierName]
						if ok {
							found = true
							fn = pair.unmarshal
						}
					})

					if !found {
						continue
					}
					if nil == fn {
						return ErrModifierNotReversible
					}

					raw, rawErr = fn(raw)
					if nil != rawErr {
						return erorr.Errorf("json: problem applying reverse modifier %q during unmarshal: %w", modifierName, rawErr)
					}
				}

				permissive := mode == unmarshalModeUnobstructed
				fieldSc := newScanner(raw, permissive)
				err = receiver.unmarshalValue(fieldSc, fieldValue, mode, path, errs)
			} else {
				err = receiver.unmarshalValue(sc, fieldValue, mode, path, errs)
			}
			if nil != err {
				return err
			}
			path.pop()
		}

	readSeparator:
		// Read comma or expect closing brace.
		tok, err = sc.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenComma {
			sc.next() // consume comma

			// Handle trailing comma in permissive mode.
			if mode == unmarshalModeUnobstructed {
				tok, err = sc.peek()
				if nil != err {
					return err
				}
				if tok.kind == tokenObjectEnd {
					sc.next() // consume }
					return nil
				}
			}
		}
	}
}

// validateConst validates a JSON value against the expected Const[T] constant.
func (receiver *Usher) validateConst(
	sc *scanner,
	dst reflect.Value,
	sf structField,
	path *jsonPath,
	errs *[]error,
) error {
	// Read the raw JSON value.
	raw, err := sc.scanRawValue()
	if nil != err {
		return err
	}

	// If there is no json.value tag, we cannot validate.
	if "" == sf.constTag {
		return nil
	}

	// Get a zero value of the Const[T] type to call DecodeFromString on.
	constVal := reflect.New(sf.fieldType).Elem().Interface()
	czer, ok := constVal.(Constantizer)
	if !ok {
		return nil
	}

	// Decode the expected value from the tag.
	expected, err := czer.DecodeFromString(sf.constTag)
	if nil != err {
		return erorr.Errorf("json: problem decoding const tag %q at %s: %w", sf.constTag, path.String(), err)
	}

	// Parse the actual JSON value to the same Go type.
	actual, err := parseJSONValueAs(raw, expected)
	if nil != err {
		// If we can't parse the actual value, it's a mismatch.
		*errs = append(*errs, ConstMismatchError{
			Path:      path.String(),
			FieldName: sf.name,
			Expected:  expected,
			Actual:    string(raw),
		})
		return nil
	}

	// Compare.
	if !reflect.DeepEqual(expected, actual) {
		*errs = append(*errs, ConstMismatchError{
			Path:      path.String(),
			FieldName: sf.name,
			Expected:  expected,
			Actual:    actual,
		})
	}

	return nil
}

// parseJSONValueAs parses raw JSON bytes into the same Go type as the exemplar value.
func parseJSONValueAs(raw []byte, exemplar any) (any, error) {
	trimmed := trim(raw)

	switch exemplar.(type) {
	case bool:
		var val bool
		if err := UnmarshalBool(trimmed, &val); nil != err {
			return nil, err
		}
		return val, nil

	case string:
		// Must be a JSON string.
		str, err := unquoteString(trimmed)
		if nil != err {
			return nil, err
		}
		return str, nil

	case int:
		return parseJSONInt[int](trimmed)
	case int8:
		return parseJSONInt[int8](trimmed)
	case int16:
		return parseJSONInt[int16](trimmed)
	case int32:
		return parseJSONInt[int32](trimmed)
	case int64:
		return parseJSONInt[int64](trimmed)

	case uint:
		return parseJSONUint[uint](trimmed)
	case uint8:
		return parseJSONUint[uint8](trimmed)
	case uint16:
		return parseJSONUint[uint16](trimmed)
	case uint32:
		return parseJSONUint[uint32](trimmed)
	case uint64:
		return parseJSONUint[uint64](trimmed)

	case float32:
		f, err := strconv.ParseFloat(string(trimmed), 32)
		if nil != err {
			return nil, err
		}
		return float32(f), nil

	case float64:
		f, err := strconv.ParseFloat(string(trimmed), 64)
		if nil != err {
			return nil, err
		}
		return f, nil

	default:
		return nil, erorr.Errorf("json: unsupported const type %T", exemplar)
	}
}
