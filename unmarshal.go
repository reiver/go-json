package json

// Unmarshal deserializes JSON data into the value pointed to by dst.
// Unknown JSON fields are silently ignored.
//
// It is equivalent to calling DefaultUsher.Unmarshal().
func Unmarshal(data []byte, dst any) error {
	return DefaultUsher.Unmarshal(data, dst)
}

// UnobstructedUnmarshal deserializes JSON data permissively.
// It allows trailing commas, comments, and leading plus-signs.
// Unknown JSON fields are silently ignored.
//
// It is equivalent to calling DefaultUsher.UnobstructedUnmarshal().
func UnobstructedUnmarshal(data []byte, dst any) error {
	return DefaultUsher.UnobstructedUnmarshal(data, dst)
}

// ObstructedUnmarshal deserializes JSON data strictly.
// Returns an error if unknown JSON fields are encountered.
//
// It is equivalent to calling DefaultUsher.ObstructedUnmarshal().
func ObstructedUnmarshal(data []byte, dst any) error {
	return DefaultUsher.ObstructedUnmarshal(data, dst)
}
