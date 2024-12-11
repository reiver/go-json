package json

// MarshalInt16 returns the JSON version of a Go int16.
func MarshalInt16(value uint16) []byte {
	return MarshalInt64(int64(value))
}
