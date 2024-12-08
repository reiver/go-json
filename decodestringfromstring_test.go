package json

import (
	"testing"

	"reflect"
)

func TestDecodeStringFromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:    "",
			Expected: "",
		},



		{
			Value:    "apple",
			Expected: "apple",
		},
		{
			Value:    "banana",
			Expected: "banana",
		},
		{
			Value:    "cherry",
			Expected: "cherry",
		},
		{
			Value:    "apple banana  cherry   ",
			Expected: "apple banana  cherry   ",
		},



		{
			Value:    "apple banana  cherry   ",
			Expected: "apple banana  cherry   ",
		},



		{
			Value:    "\"",
			Expected: "\"",
		},
		{
			Value:    "\\",
			Expected: "\\",
		},
		{
			Value:    "\n",
			Expected: "\n",
		},
		{
			Value:    "\t",
			Expected: "\t",
		},
		{
			Value:    "\v",
			Expected: "\v",
		},
		{
			Value:    "\b",
			Expected: "\b",
		},
		{
			Value:    "\r",
			Expected: "\r",
		},
		{
			Value:    "\f",
			Expected: "\f",
		},
		{
			Value:    "\x20",
			Expected: "\x20",
		},
		{
			Value:    "\U0001F642",
			Expected: "\U0001F642",
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeStringFromString(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual decoded value is not what wad expected", testNumber)
			t.Logf("EXPECTED: (%T) %#v", expected, expected)
			t.Logf("ACTUAL:   (%T) %#v", actual, actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
