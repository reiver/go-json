package json

import (
	"testing"

	"reflect"
)

func TestDecodeBoolFromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:   "false",
			Expected: false,
		},
		{
			Value:   "true",
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeBoolFromString(test.Value)
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
