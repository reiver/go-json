package json

var (
	// DefaultUsher is the Usher that the json.Marshal() function uses.
	//
	// Implanting a modifier into DefaultUsher into DefaultUsher modifies how the json.Marshal() behaves.
	DefaultUsher Usher
)

func init() {
	DefaultUsher.ImplantModifier("string", StringModifierFunc)
	DefaultUsher.ImplantModifier("bare", BareModifierFunc)
}
