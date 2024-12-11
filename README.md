# go-json

Package **json** implements JSON encoders and decoders, for the Go programming language.

Package **json** is meant to be a replacement for the Go built-in `"encoding/json"` package.

Package **json** also includes a number of addtional useful features (that the Go built-in `"encoding/json"` package does _not_ have).

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

## Installation

To install package **json** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-json
```

## Author

Package **json** was written by [Charles Iliya Krempeaux](http://reiver.link)
