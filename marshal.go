package json

// Marshal returns the JSON version of 'value'.
//
// See [Usher.Marshal] for more information.
func Marshal(value any) ([]byte, error) {
	return DefaultUsher.Marshal(value)
}
