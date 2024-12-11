package json

// MarshalInt returns the JSON version of a Go int.
func MarshalInt(value int) []byte {
	return MarshalInt64(int64(value))
}
