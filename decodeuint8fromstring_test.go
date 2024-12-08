package json

import (
	"testing"

	"reflect"
)

func TestDecodeUint8FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:         "0",
			Expected: uint8(0),
		},
		{
			Value:         "1",
			Expected: uint8(1),
		},

		{
			Value:         "127",
			Expected: uint8(127),
		},
		{
			Value:         "128",
			Expected: uint8(128),
		},

		{
			Value:         "254",
			Expected: uint8(254),
		},
		{
			Value:         "255",
			Expected: uint8(255),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeUint8FromString(test.Value)
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
