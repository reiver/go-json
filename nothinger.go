package json

// Nothinger is similar to [Emptier], but is more commonly used with optional-types (i.e., optiona-types).
//
// I.e., an optional-type has "something" in it or "nothing" in it.
// (Some optional-types call this "some" and "none".)
// A optional-type might have a method IsNothing()bool to communicate whether it has "nothing" or "something" in it.
//
// If the optional-type has such a IsNothing()bool method, then [Marshal] and [Usher.Marshal] make use of that method for the purposes of `omitempty`.
//
// If you type in not an optional-type, then it should probably instead implement [Emptier].
type Nothinger interface {
	IsNothing() bool
}
