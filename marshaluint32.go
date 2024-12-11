package json

// MarshalUint64 returns the JSON version of a Go uint64.
func MarshalUint32(value uint32) []byte {
	return MarshalUint64(uint64(value))
}
