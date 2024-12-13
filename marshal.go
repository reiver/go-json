package json

// Marshal returns the JSON version of 'value'.
//
// For example:
//
//	bytes, err := json.Marshal(value)
//
// See [Usher.Marshal] for more information.
//
// Calling:
//
//	bytes, err := json.Marshal(value)
//
// Is the same as calling:
//
//	bytes, err := json.DefaultUsher.Marshal(value)
//
// json.Marshal() is a convenience function for json.DefaultUsher.Marshal().
func Marshal(value any) ([]byte, error) {
	return DefaultUsher.Marshal(value)
}
