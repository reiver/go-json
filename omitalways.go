package json

// A field in a struct and a value in a map that fits this interface (by having the method JSONOmitAlways()) will always be omitted from the result JSON.
type OmitAlways interface {
	JSONOmitAlways()
}
