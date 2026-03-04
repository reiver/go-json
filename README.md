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

## Custom Modifiers

Package **json** supports _custom modifiers_ through struct tags.
A modifier transforms the already-marshaled JSON bytes of a struct field.

### Tag Format

```golang
Field Type `json:"name,modifier1,modifier2,..."`
```

Any struct tag option that isn't `omitempty` or `-` is treated as a modifier name.
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

To add a custom modifier to the global `DefaultUsher` (so it works with `json.Marshal`):

```golang
json.DefaultUsher.ImplantModifier("mymod", myModifierFunc)
```

## Installation

To install package **json** do the following:
```
GOPROXY=direct go get github.com/reiver/go-json
```

## Author

Package **json** was written by [Charles Iliya Krempeaux](http://reiver.link)
