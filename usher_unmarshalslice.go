package json

import (
	"reflect"

	"codeberg.org/reiver/go-erorr"
)

// unmarshalSlice unmarshals a JSON array into a Go slice.
// The opening '[' has NOT yet been consumed.
func (receiver *Usher) unmarshalSlice(
	sc *scanner,
	dst reflect.Value,
	mode unmarshalMode,
	path *jsonPath,
	errs *[]error,
) error {
	// Consume the opening bracket.
	tok, err := sc.next()
	if nil != err {
		return err
	}
	if tok.kind != tokenArrayBegin {
		return erorr.Errorf("json: expected '[' at position %d", tok.pos)
	}

	elemType := dst.Type().Elem()
	slice := reflect.MakeSlice(dst.Type(), 0, 0)

	index := 0
	for {
		// Check for closing bracket.
		tok, err := sc.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenArrayEnd {
			sc.next() // consume ]
			dst.Set(slice)
			return nil
		}

		// Read the element.
		path.pushIndex(index)
		elem := reflect.New(elemType).Elem()
		err = receiver.unmarshalValue(sc, elem, mode, path, errs)
		if nil != err {
			return err
		}
		path.pop()

		slice = reflect.Append(slice, elem)
		index++

		// Comma or closing bracket.
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
				if tok.kind == tokenArrayEnd {
					sc.next() // consume ]
					dst.Set(slice)
					return nil
				}
			}
		}
	}
}

// unmarshalArray unmarshals a JSON array into a Go array.
// The opening '[' has NOT yet been consumed.
func (receiver *Usher) unmarshalArray(
	sc *scanner,
	dst reflect.Value,
	mode unmarshalMode,
	path *jsonPath,
	errs *[]error,
) error {
	// Consume the opening bracket.
	tok, err := sc.next()
	if nil != err {
		return err
	}
	if tok.kind != tokenArrayBegin {
		return erorr.Errorf("json: expected '[' at position %d", tok.pos)
	}

	arrayLen := dst.Len()
	index := 0

	for {
		// Check for closing bracket.
		tok, err := sc.peek()
		if nil != err {
			return err
		}
		if tok.kind == tokenArrayEnd {
			sc.next() // consume ]
			return nil
		}

		// Read the element.
		path.pushIndex(index)

		if index < arrayLen {
			err = receiver.unmarshalValue(sc, dst.Index(index), mode, path, errs)
		} else {
			// Excess elements: skip them.
			err = sc.skipValue()
		}
		if nil != err {
			return err
		}
		path.pop()

		index++

		// Comma or closing bracket.
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
				if tok.kind == tokenArrayEnd {
					sc.next() // consume ]
					return nil
				}
			}
		}
	}
}
