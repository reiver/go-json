package json

// Emptier works with the `omitempty` struct-tag, recognized by [Marshal] and [Usher.Marshal].
//
// For example, a custom type might look like:
//
//	type MyType struct {
//		// ...
//	}
//	
//	func (receiver MyType) IsEmpty() bool {
//		// ...
//	}
//
// And it might be used similar to the following:
//
//	var MyStruct struct {
//		Apple  string `json:"apple,omitempty"`
//		Banana MyType `json:"banana,omitempty"` // <---------
//		Cherry int    `json:"cherry"`
//	}
//	
//	// ...
//
//	var value MyStruct // = ...
//	
//	bytes, err := json.Marshal(value)
//
// An alternative to Emptier is [Nothinger], which is more commonly used with optional-types (i.e., optiona-types).
type Emptier interface {
	IsEmpty() bool
}
