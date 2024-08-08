package json

// Marshal return the JSON version of 'value'.
func Marshal(value any) ([]byte, error) {
	return DefaultUsher.Marshal(value)
}
