package json

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrModifierNotReversible = erorr.Error("json: modifier is not reversible")
	ErrNilReceiver           = erorr.Error("json: nil receiver")
	ErrNotBool               = erorr.Error("json: not bool")
	ErrNotJSONNumber         = erorr.Error("json: not JSON number")
)

const (
	errBadReflection                      = erorr.Error("json: bad reflection")
	errNoBytes                            = erorr.Error("json: no bytes")
	errNotStringMissingBeginQuotationMark = erorr.Error("json: not string — missing begin quotation-mark")
	errNotStringMissingEndQuotationMark   = erorr.Error("json: not string — missing end quotation-mark")
	errNotStruct                          = erorr.Error("json: not struct")

	errNilDestination      = erorr.Error("json: nil destination")
	errNotPointer          = erorr.Error("json: destination is not a pointer")
	errUnexpectedToken     = erorr.Error("json: unexpected token")
	errUnterminatedString  = erorr.Error("json: unterminated string")
	errUnterminatedComment = erorr.Error("json: unterminated comment")
	errTrailingData        = erorr.Error("json: trailing data after value")
	errInvalidEscape       = erorr.Error("json: invalid escape sequence")
)
