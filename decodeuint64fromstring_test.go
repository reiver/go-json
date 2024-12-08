package json

import (
	"testing"

	"reflect"
)

func TestDecodeUint64FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:          "0",
			Expected: uint64(0),
		},
		{
			Value:          "1",
			Expected: uint64(1),
		},

		{
			Value:          "127",
			Expected: uint64(127),
		},
		{
			Value:          "128",
			Expected: uint64(128),
		},

		{
			Value:          "254",
			Expected: uint64(254),
		},
		{
			Value:          "255",
			Expected: uint64(255),
		},
		{
			Value:          "256",
			Expected: uint64(256),
		},
		{
			Value:          "257",
			Expected: uint64(257),
		},

		{
			Value:          "32767",
			Expected: uint64(32767),
		},
		{
			Value:          "32768",
			Expected: uint64(32768),
		},
		{
			Value:          "32769",
			Expected: uint64(32769),
		},

		{
			Value:          "65535",
			Expected: uint64(65535),
		},
		{
			Value:          "65536",
			Expected: uint64(65536),
		},
		{
			Value:          "65537",
			Expected: uint64(65537),
		},

		{
			Value:          "2147483647",
			Expected: uint64(2147483647),
		},
		{
			Value:          "2147483648",
			Expected: uint64(2147483648),
		},
		{
			Value:          "2147483649",
			Expected: uint64(2147483649),
		},

		{
			Value:          "4294967295",
			Expected: uint64(4294967295),
		},
		{
			Value:          "4294967296",
			Expected: uint64(4294967296),
		},
		{
			Value:          "4294967297",
			Expected: uint64(4294967297),
		},

		{
			Value:          "18446744073709551615",
			Expected: uint64(18446744073709551615),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeUint64FromString(test.Value)
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
