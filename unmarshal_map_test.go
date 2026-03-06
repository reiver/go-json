package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnmarshal_map(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes: []byte(`{"apple":"red","banana":"yellow"}`),
			Dst: func() any {
				return &map[string]string{}
			},
			Expected: map[string]string{
				"apple":  "red",
				"banana": "yellow",
			},
		},



		{
			Bytes: []byte(`{"a":1,"b":2,"c":3}`),
			Dst: func() any {
				return &map[string]int{}
			},
			Expected: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},



		{
			Bytes: []byte(`{"mixed":"hello","num":42,"flag":true,"nothing":null}`),
			Dst: func() any {
				return &map[string]any{}
			},
			Expected: map[string]any{
				"mixed":   "hello",
				"num":     json.MustParseNumberString("42"),
				"flag":    true,
				"nothing": nil,
			},
		},



		{
			Bytes: []byte(`{}`),
			Dst: func() any {
				return &map[string]string{}
			},
			Expected: map[string]string{},
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.Unmarshal(test.Bytes, dst)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		actual := reflect.ValueOf(dst).Elem().Interface()

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", test.Expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}
	}
}
