package json

// MarshalBool returns the JSON version of a Go bool.
func MarshalBool(value bool) []byte {
	var f [5]byte = [5]byte{'f','a','l','s','e'}
	var t [4]byte = [4]byte{'t','r','u','e'}

	switch value {
	case false:
		return f[:]
	default:
		return t[:]
	}
}
