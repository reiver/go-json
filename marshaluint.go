package json

// MarshalUint returns the JSON version of a Go uint.
func MarshalUint(value uint) []byte {
	return MarshalUint64(uint64(value))
}
