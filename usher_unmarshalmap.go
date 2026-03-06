package json

import (
	"reflect"

	"codeberg.org/reiver/go-erorr"
)

// unmarshalMap unmarshals a JSON object into a Go map.
// The opening '{' has NOT yet been consumed.
func (receiver *Usher) unmarshalMap(
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

	// Initialize the map if nil.
	if dst.IsNil() {
		dst.Set(reflect.MakeMap(dst.Type()))
	}

	keyType := dst.Type().Key()
	valType := dst.Type().Elem()

	for {
		// Check for closing brace.
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
		keyStr, err := unquoteString(keyTok.value)
		if nil != err {
			return err
		}

		// Convert key to the map's key type.
		mapKey := reflect.New(keyType).Elem()
		switch keyType.Kind() {
		case reflect.String:
			mapKey.SetString(keyStr)
		default:
			return erorr.Errorf("json: unsupported map key type %s at %s", keyType, path.String())
		}

		// Read the colon.
		colonTok, err := sc.next()
		if nil != err {
			return err
		}
		if colonTok.kind != tokenColon {
			return erorr.Errorf("json: expected colon at position %d", colonTok.pos)
		}

		// Read the value.
		path.pushKey(keyStr)
		mapVal := reflect.New(valType).Elem()
		err = receiver.unmarshalValue(sc, mapVal, mode, path, errs)
		if nil != err {
			return err
		}
		path.pop()

		dst.SetMapIndex(mapKey, mapVal)

		// Comma or closing brace.
		tok, err = sc.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenComma {
			sc.next()

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
