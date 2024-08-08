package json

// Marshaler is something that can marshal itself into JSON.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
