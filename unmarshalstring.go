package json

// UnmarshalString JSON-unmarshals a JSON string into a Go string.
func UnmarshalString(data []byte, dst *string) error {
	str, err := unquoteString(data)
	if nil != err {
		return err
	}

	*dst = str
	return nil
}
