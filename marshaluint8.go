package json

// MarshalUint8 returns the JSON version of a Go uint64.
func MarshalUint8(value uint8) []byte {
	return MarshalUint64(uint64(value))
}
