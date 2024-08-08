package json_test

import (
	"github.com/reiver/go-opt"
)

type DemoType2 struct {
        Apple  opt.Optional[string] `json:"apple,omitempty"`
        Banana opt.Optional[string] `json:"banana,omitempty"`
        Cherry opt.Optional[string] `json:"cherry,omitempty"`
}
