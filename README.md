# go-json

Package **json** implements JSON encoders and decoders, for the Go programming-language (golang).

Package **json** is meant to be a replacement for the Go built-in `"encoding/json"` package.

Package **json** also includes a number of addtional useful features (that the Go built-in `"encoding/json"` package does _not_ have).

Package **json** is a custom JSON encoding/decoding library for the Go programming-language (golang), designed as a replacement for the Go built-in`"encoding/json"` package with additional features like _custom modifiers_, _constant fields_, and _extended struct tag support_.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-json

[![GoDoc](https://godoc.org/github.com/reiver/go-json?status.svg)](https://godoc.org/github.com/reiver/go-json)

## Examples

To marshal a Go value to JSON, you can do something similar to the following:

```golang
import "github.com/reiver/go-json"

// ...

jsonBytes, err := json.Marshal(value)
```

To unmarshal a JSON value to Go, you can do something similar to the following:

```golang
import "github.com/reiver/go-json"

// ...

err := json.Unarshal(jsonBytes, &value)
```

## Custom Modifiers

Package **json** supports _custom modifiers_ through struct tags.
A modifier transforms the already-marshaled JSON bytes of a struct field.

### Tag Format

```golang
Field Type `json:"name,modifier1,modifier2,..."`
```

Any struct tag option that isn't `omitempty`, `nullempty`, or `-` is treated as a modifier name.
Multiple modifiers are applied in order, left to right.

### Built-in Modifiers

The `DefaultUsher` (used by `json.Marshal`) comes with two built-in modifiers:

**`string`** — wraps the marshaled value in quotes:

```golang
type MyStruct struct {
	Banana int `json:"banana,string"`
}

// Banana with value 5 produces: {"banana":"5"}
```

**`bare`** — strips quotes from a JSON string, outputting the raw content:

```golang
type MyStruct struct {
	Banana string `json:"banana,bare"`
}

// Banana with value "true" produces: {"banana":true}
```

### Registering Custom Modifiers

You can register your own modifiers on an `Usher` instance:

```golang
var usher json.Usher

usher.ImplantModifier("digest", func(data []byte) ([]byte, error) {
	// transform the marshaled bytes
	// ...
	return transformed, nil
})

jsonBytes, err := usher.Marshal(value)
```

A `ModifierFunc` has the signature `func([]byte) ([]byte, error)`.
It receives the marshaled JSON bytes for a field and returns the transformed bytes.

Modifiers can be chained:

```golang
Field string `json:"field,omitempty,string,mymod"`
```

This applies `string` first, then `mymod`.
If any modifier returns an `ErrorEmpty` error and the field has `omitempty`, the field is omitted from the output.
If the field has `nullempty` instead, the field outputs `null`.

To add a custom modifier to the global `DefaultUsher` (so it works with `json.Marshal`):

```golang
json.DefaultUsher.ImplantModifier("mymod", myModifierFunc)
```

## Null Empty (`nullempty`)

The `nullempty` struct tag option outputs JSON `null` for a field when its value is empty, instead of omitting the field entirely (which is what `omitempty` does).

```golang
type MyStruct struct {
	Name  string   `json:"name"`
	Value string   `json:"value,nullempty"`
	Items []string `json:"items,nullempty"`
}

// With zero values produces: {"name":"","value":null,"items":null}
```

Emptiness is determined using the same checks as `omitempty`:
- Types implementing `Emptier` (`IsEmpty() bool`)
- Types implementing `Nothinger` (`IsNothing() bool`)
- Zero values (via `reflect.DeepEqual`)
- Empty slices, maps, and arrays (length 0)

When a type implements `Emptier` or `Nothinger`, that interface is trusted and the fallback zero-value/length checks are not applied.

If a `MarshalJSON()` method or a modifier returns an `ErrorEmpty` error and the field has `nullempty`, the field outputs `null` instead of being omitted.

## Installation

To install package **json** do the following:
```
GOPROXY=direct go get github.com/reiver/go-json
```

## Author

Package **json** was written by [Charles Iliya Krempeaux](http://reiver.link)
