package json

// ModifierFunc is the type of a modifer.
//
// Modifers as used to modify the the a marshaled struct-field.
type ModifierFunc func([]byte) ([]byte, error)
