package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_array(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		// Basic arrays of primitive types.
		{
			Value: [0]string{},
			Expected: `[]`,
		},
		{
			Value: [1]string{"once"},
			Expected: `["once"]`,
		},
		{
			Value: [2]string{"once","twice"},
			Expected: `["once","twice"]`,
		},
		{
			Value: [3]string{"once","twice","thrice"},
			Expected: `["once","twice","thrice"]`,
		},
		{
			Value: [4]string{"once","twice","thrice","fource"},
			Expected: `["once","twice","thrice","fource"]`,
		},
		{
			Value: [3]int{1,2,3},
			Expected: `[1,2,3]`,
		},
		{
			Value: [2]bool{true,false},
			Expected: `[true,false]`,
		},



		// Array of structs with Const fields — demonstrates that
		// arrays should go through the custom pipeline, not gojson.Marshal.
		// gojson.Marshal does not understand Const[T] or the json.value tag,
		// so it would marshal Const[T] as {} (empty struct).
		{
			Value: [2]struct{
				Name    string             `json:"name"`
				Country json.Const[string] `json:"country" json.value:"Canada"`
			}{
				{Name: "Alice"},
				{Name: "Bob"},
			},
			Expected: `[{"name":"Alice","country":"Canada"},{"name":"Bob","country":"Canada"}]`,
		},



		// Array as a struct field.
		{
			Value: struct{
				Items [3]string `json:"items"`
			}{
				Items: [3]string{"a","b","c"},
			},
			Expected: `{"items":["a","b","c"]}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual json-marshaled value for the %T is not what was expected.", testNumber, test.Value)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				t.Logf("VALUE-KIND: %s", reflect.TypeOf(test.Value).Kind())
				t.Logf("VALUE-TYPE: %T", test.Value)
				continue
			}
		}
	}
}
