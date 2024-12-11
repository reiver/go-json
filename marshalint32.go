package json

// MarshalInt32 returns the JSON version of a Go int32.
func MarshalInt32(value uint32) []byte {
	return MarshalInt64(int64(value))
}
