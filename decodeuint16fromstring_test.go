package json

import (
	"testing"

	"reflect"
)

func TestDecodeUint16FromString(t *testing.T) {

	tests := []struct{
		Value string
		Expected any
	}{
		{
			Value:          "0",
			Expected: uint16(0),
		},
		{
			Value:          "1",
			Expected: uint16(1),
		},

		{
			Value:          "127",
			Expected: uint16(127),
		},
		{
			Value:          "128",
			Expected: uint16(128),
		},

		{
			Value:          "254",
			Expected: uint16(254),
		},
		{
			Value:          "255",
			Expected: uint16(255),
		},
		{
			Value:          "256",
			Expected: uint16(256),
		},
		{
			Value:          "257",
			Expected: uint16(257),
		},

		{
			Value:          "32767",
			Expected: uint16(32767),
		},
		{
			Value:          "32768",
			Expected: uint16(32768),
		},
		{
			Value:          "32769",
			Expected: uint16(32769),
		},

		{
			Value:          "65535",
			Expected: uint16(65535),
		},
	}

	for testNumber, test := range tests {

		actual, err := decodeUint16FromString(test.Value)
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
