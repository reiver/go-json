package json

import (
	"testing"

	"reflect"
)

func TestDecodeInt16FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:         "-32768",
			Expected: int16(-32768),
		},
		{
			Value:         "-32767",
			Expected: int16(-32767),
		},

		{
			Value:         "-128",
			Expected: int16(-128),
		},
		{
			Value:         "-127",
			Expected: int16(-127),
		},

		{
			Value:         "-2",
			Expected: int16(-2),
		},
		{
			Value:         "-1",
			Expected: int16(-1),
		},
		{
			Value:         "0",
			Expected: int16(0),
		},
		{
			Value:         "1",
			Expected: int16(1),
		},

		{
			Value:         "127",
			Expected: int16(127),
		},
		{
			Value:         "128",
			Expected: int16(128),
		},

		{
			Value:         "32767",
			Expected: int16(32767),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeInt16FromString(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("Value:  %q", test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual decoded value is not what wad expected", testNumber)
			t.Logf("EXPECTED: (%T) %#v", expected, expected)
			t.Logf("ACTUAL:   (%T) %#v", actual, actual)
			t.Logf("Value:  %q", test.Value)
			continue
		}
	}
}
