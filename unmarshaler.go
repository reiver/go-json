package json

// Unmarshaler is something that can unmarshal itself from JSON.
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
