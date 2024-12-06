package json

import (
	"github.com/reiver/go-lck"
)

// Usher marshals a Go type into JSON.
//
// If you want the "string" modifier you will need to call json.Usher.ImplantModifier() to add it:
//
//	var jsonUsher json.Usher
//
//	jsonUsher.ImplantModifier("string", json.StringModifierFunc)
//
// You can also add your own modifiers:
//
//	var jsonUsher json.Usher
//
//	jsonUsher.ImplantModifier("digest", digestFunc)
type Usher struct {
	modifierFuncs lck.Lockable[map[string]ModifierFunc]
}

func (receiver *Usher) ImplantModifier(name string, fn ModifierFunc) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.modifierFuncs.Let(func(modifiers *map[string]ModifierFunc){
		if nil == *modifiers {
			*modifiers = map[string]ModifierFunc{}
		}

		switch fn {
		case nil:
			delete(*modifiers, name)
		default:
			(*modifiers)[name] = fn
		}
	})
}
