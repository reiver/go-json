package json

import (
	"github.com/reiver/go-erorr"
)

// BareModifierFunc is the modifier that, in the DefaultUsher, is the implementation behind the "bare" modifier.
//
// For example:
//
//	struct {
//
//		// ...
//
//		Banana string `json:"banana,bare"`
//
//		// ...
//
//	}
func BareModifierFunc(bytes []byte) ([]byte, error) {
	var length int = len(bytes)
	if length <= 0 {
		return nil, erorr.Errorf("json: bare modifier-func encounter a problem: %w", errNoBytes)
	}

	{
		var byte0 byte = bytes[0]
		if '"' != byte0 {
			return nil, erorr.Errorf("json: bare modifier-func encounter a problem: %w", errNotStringMissingBeginQuotationMark)
		}
	}

	{
		var byteLast byte = bytes[length-1]
		if '"' != byteLast {
			return nil, erorr.Errorf("json: bare modifier-func encounter a problem: %w", errNotStringMissingEndQuotationMark)
		}
	}

	var str string
	{
		err := UnmarshalString(bytes, &str)
		if nil != err {
			return nil, erorr.Errorf("json: bare modifier-func encounter a problem: %w", err)
		}
	}

//@TODO: should the value of 'str' be validated as a JSON value?

	return []byte(str), nil
}
