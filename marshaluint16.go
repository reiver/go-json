package json

// MarshalUint64 returns the JSON version of a Go uint64.
func MarshalUint16(value uint16) []byte {
	return MarshalUint64(uint64(value))
}
