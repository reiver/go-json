package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_uint8(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: uint8(0),
			Expected:    "0",
		},
		{
			Value: uint8(1),
			Expected:    "1",
		},
		{
			Value: uint8(2),
			Expected:    "2",
		},
		{
			Value: uint8(3),
			Expected:    "3",
		},
		{
			Value: uint8(4),
			Expected:    "4",
		},
		{
			Value: uint8(5),
			Expected:    "5",
		},

		{
			Value: uint8(12),
			Expected:    "12",
		},

		{
			Value: uint8(100),
			Expected:    "100",
		},

		{
			Value: uint8(123),
			Expected:    "123",
		},

		{
			Value: uint8(126),
			Expected:    "126",
		},
		{
			Value: uint8(127),
			Expected:    "127",
		},
		{
			Value: uint8(128),
			Expected:    "128",
		},
		{
			Value: uint8(129),
			Expected:    "129",
		},

		{
			Value: uint8(200),
			Expected:    "200",
		},

		{
			Value: uint8(254),
			Expected:    "254",
		},
		{
			Value: uint8(255),
			Expected:    "255",
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
