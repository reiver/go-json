package json

import (
	"unsafe"
)

// StringModifierFunc is the modifier that, in the DefaultUsher, is the implementation behind the "string" modifier.
//
// For example:
//
//	struct {
//
//		// ...
//
//		Banana int `json:"banana,string"`
//
//		// ...
//
//	}
func StringModifierFunc(bytes []byte) ([]byte, error) {
	var length int = len(bytes)
	var ptr *byte = unsafe.SliceData(bytes)

	var str string = unsafe.String(ptr, length)

	return Marshal(str)
}
