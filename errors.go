package json

import (
	"github.com/reiver/go-erorr"
)

const (
	errBadReflection                      = erorr.Error("json: bad reflection")
	errNilReceiver                        = erorr.Error("json: nil receiver")
	errNoBytes                            = erorr.Error("json: no bytes")
	errNotStringMissingBeginQuotationMark = erorr.Error("json: not string — missing begin quotation-mark")
	errNotStringMissingEndQuotationMark   = erorr.Error("json: not string — missing end quotation-mark")
	errNotStruct                          = erorr.Error("json: not struct")
)
