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
//	jsonUsher.ImplantModifier("string", json.StringModifierFunc, json.BareModifierFunc)
//
// (Note that [BareModifierFunc] is also included because it is the inverse-operation to [StringModifierFunc].)
//
// You can also add your own modifiers:
//
//	var jsonUsher json.Usher
//
//	jsonUsher.ImplantModifier("digest", digestFunc, nil)
type Usher struct {
	modifierFuncs lck.Lockable[map[string]modifierPair]
}

func (receiver *Usher) ImplantModifier(name string, marshalFn ModifierFunc, unmarshalFn ModifierFunc) {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	receiver.modifierFuncs.Let(func(modifiers *map[string]modifierPair){
		if nil == *modifiers {
			*modifiers = map[string]modifierPair{}
		}

		if nil == marshalFn && nil == unmarshalFn {
			delete(*modifiers, name)
		} else {
			(*modifiers)[name] = modifierPair{
				marshal:   marshalFn,
				unmarshal: unmarshalFn,
			}
		}
	})
}
