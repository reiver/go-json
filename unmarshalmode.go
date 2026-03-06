package json

// unmarshalMode controls the behavior of the unmarshaling engine.
type unmarshalMode int

const (
	// unmarshalModeStandard ignores unknown fields, requires standard JSON syntax.
	unmarshalModeStandard unmarshalMode = iota

	// unmarshalModeUnobstructed is permissive: trailing commas, comments, ignores unknown fields.
	unmarshalModeUnobstructed

	// unmarshalModeObstructed is strict: errors on unknown fields, requires standard JSON syntax.
	unmarshalModeObstructed
)
