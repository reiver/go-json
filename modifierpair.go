package json

type modifierPair struct {
	marshal   ModifierFunc
	unmarshal ModifierFunc
}
