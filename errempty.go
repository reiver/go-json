package json

import (
	"strings"
)

// ErrEmpty returns an error of type [ErrorEmpty].
//
// [ErrorEmpty] is used with the "omitempty" struct-field tag option.
func ErrEmpty(message string) error {
	return internalErrorEmpty{
		message:message,
	}
}

// ErrorEmpty is a special type of error.
//
// ErrorEmpty is used with the "omitempty" struct-field tag option.
//
// If a (custom) type's MarshalJSON() function returns an error of type ErrorEmpty, and
// a field of that type has the "omitempty" struct-field tag option then, that field will
// be omitted in the marshaled-JSON.
//
// You can, of course, create your own (custom) error type that fits this ErrorEmpty interface, or
// you can use the [ErrEmpty] to create that ErrorEmpty error for you.
type ErrorEmpty interface {
	error
	ErrorEmpty()
}

type internalErrorEmpty struct {
	message string
}

var _ ErrorEmpty = internalErrorEmpty{}

func (receiver internalErrorEmpty) Error() string {
	var builder strings.Builder

	builder.WriteString("empty")

	{
		var message string = receiver.message

		if "" != message {
			builder.WriteString(" — ")
			builder.WriteString(message)
		}
	}

	return builder.String()
}

func (internalErrorEmpty) ErrorEmpty() {
	// nothing here
}
