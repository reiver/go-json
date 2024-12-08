package json

import (
	"testing"

	"reflect"
)

func TestDecodeInt64FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:         "-2147483649",
			Expected: int64(-2147483649),
		},
		{
			Value:         "-2147483648",
			Expected: int64(-2147483648),
		},
		{
			Value:         "-2147483647",
			Expected: int64(-2147483647),
		},

		{
			Value:         "-32769",
			Expected: int64(-32769),
		},
		{
			Value:         "-32768",
			Expected: int64(-32768),
		},
		{
			Value:         "-32767",
			Expected: int64(-32767),
		},

		{
			Value:         "-128",
			Expected: int64(-128),
		},
		{
			Value:         "-127",
			Expected: int64(-127),
		},

		{
			Value:         "-2",
			Expected: int64(-2),
		},
		{
			Value:         "-1",
			Expected: int64(-1),
		},
		{
			Value:         "0",
			Expected: int64(0),
		},
		{
			Value:         "1",
			Expected: int64(1),
		},

		{
			Value:         "127",
			Expected: int64(127),
		},
		{
			Value:         "128",
			Expected: int64(128),
		},

		{
			Value:         "32767",
			Expected: int64(32767),
		},
		{
			Value:         "32768",
			Expected: int64(32768),
		},
		{
			Value:         "32769",
			Expected: int64(32769),
		},

		{
			Value:         "2147483647",
			Expected: int64(2147483647),
		},
		{
			Value:         "2147483648",
			Expected: int64(2147483648),
		},
		{
			Value:         "2147483649",
			Expected: int64(2147483649),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeInt64FromString(test.Value)
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
