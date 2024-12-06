package json

// Marshal return the JSON version of 'value'.
func MergeAndMarshal(values ...any) ([]byte, error) {
	return DefaultUsher.MergeAndMarshal(values...)
}
