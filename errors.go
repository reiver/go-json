package json

import (
	"github.com/reiver/go-erorr"
)

const (
	errBadReflection = erorr.Error("json: bad reflection")
	errNilReceiver   = erorr.Error("json: nil receiver")
	errNotStruct     = erorr.Error("json: not struct")
)
