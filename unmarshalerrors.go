package json

import (
	"fmt"
	"strings"
)

// ConstMismatchError is returned when a JSON value does not match
// the expected Const[T] value declared in the json.value struct tag.
type ConstMismatchError struct {
	Path      string
	FieldName string
	Expected  any
	Actual    any
}

func (e ConstMismatchError) Error() string {
	return fmt.Sprintf("json: const mismatch at %s (field %q): expected %v, got %v", e.Path, e.FieldName, e.Expected, e.Actual)
}

// UnknownFieldError is returned in obstructed mode when a JSON key
// does not match any struct field.
type UnknownFieldError struct {
	Path string
	Key  string
}

func (e UnknownFieldError) Error() string {
	return fmt.Sprintf("json: unknown field %q at %s", e.Key, e.Path)
}

// UnmarshalErrors collects multiple non-fatal errors from a single
// Unmarshal operation (e.g., multiple Const[T] mismatches).
type UnmarshalErrors struct {
	Errors []error
}

func (e UnmarshalErrors) Error() string {
	var builder strings.Builder

	builder.WriteString("json: unmarshal encountered ")
	builder.WriteString(fmt.Sprintf("%d", len(e.Errors)))
	builder.WriteString(" error(s):")

	for _, err := range e.Errors {
		builder.WriteString("\n\t")
		builder.WriteString(err.Error())
	}

	return builder.String()
}

func (e UnmarshalErrors) Unwrap() []error {
	return e.Errors
}
