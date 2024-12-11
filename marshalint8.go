package json

// MarshalInt8 returns the JSON version of a Go int8.
func MarshalInt8(value int8) []byte {
	return MarshalInt64(int64(value))
}
