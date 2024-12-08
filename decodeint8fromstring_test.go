package json

import (
	"testing"

	"reflect"
)

func TestDecodeInt8FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:        "-128",
			Expected: int8(-128),
		},
		{
			Value:        "-127",
			Expected: int8(-127),
		},

		{
			Value:        "-2",
			Expected: int8(-2),
		},
		{
			Value:        "-1",
			Expected: int8(-1),
		},
		{
			Value:        "0",
			Expected: int8(0),
		},
		{
			Value:        "1",
			Expected: int8(1),
		},

		{
			Value:        "127",
			Expected: int8(127),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeInt8FromString(test.Value)
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
