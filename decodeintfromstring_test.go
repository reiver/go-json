package json

import (
	"testing"

	"reflect"
)

func TestDecodeIntFromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:       "-2147483649",
			Expected: int(-2147483649),
		},
		{
			Value:       "-2147483648",
			Expected: int(-2147483648),
		},
		{
			Value:       "-2147483647",
			Expected: int(-2147483647),
		},

		{
			Value:       "-32769",
			Expected: int(-32769),
		},
		{
			Value:       "-32768",
			Expected: int(-32768),
		},
		{
			Value:       "-32767",
			Expected: int(-32767),
		},

		{
			Value:       "-128",
			Expected: int(-128),
		},
		{
			Value:       "-127",
			Expected: int(-127),
		},

		{
			Value:       "-2",
			Expected: int(-2),
		},
		{
			Value:       "-1",
			Expected: int(-1),
		},
		{
			Value:       "0",
			Expected: int(0),
		},
		{
			Value:       "1",
			Expected: int(1),
		},

		{
			Value:       "127",
			Expected: int(127),
		},
		{
			Value:       "128",
			Expected: int(128),
		},

		{
			Value:       "32767",
			Expected: int(32767),
		},
		{
			Value:       "32768",
			Expected: int(32768),
		},
		{
			Value:       "32769",
			Expected: int(32769),
		},

		{
			Value:       "2147483647",
			Expected: int(2147483647),
		},
		{
			Value:       "2147483648",
			Expected: int(2147483648),
		},
		{
			Value:       "2147483649",
			Expected: int(2147483649),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeIntFromString(test.Value)
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
