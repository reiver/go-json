package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_int8(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: int8(-128),
			Expected:   "-128",
		},
		{
			Value: int8(-127),
			Expected:   "-127",
		},
		{
			Value: int8(-126),
			Expected:   "-126",
		},

		{
			Value: int8(-123),
			Expected:   "-123",
		},

		{
			Value: int8(-100),
			Expected:   "-100",
		},

		{
			Value: int8(-12),
			Expected:   "-12",
		},

		{
			Value: int8(-5),
			Expected:   "-5",
		},
		{
			Value: int8(-4),
			Expected:   "-4",
		},
		{
			Value: int8(-3),
			Expected:   "-3",
		},
		{
			Value: int8(-2),
			Expected:   "-2",
		},
		{
			Value: int8(-1),
			Expected:   "-1",
		},
		{
			Value: int8(0),
			Expected:   "0",
		},
		{
			Value: int8(1),
			Expected:   "1",
		},
		{
			Value: int8(2),
			Expected:   "2",
		},
		{
			Value: int8(3),
			Expected:   "3",
		},
		{
			Value: int8(4),
			Expected:   "4",
		},
		{
			Value: int8(5),
			Expected:   "5",
		},

		{
			Value: int8(12),
			Expected:   "12",
		},

		{
			Value: int8(100),
			Expected:   "100",
		},

		{
			Value: int8(123),
			Expected:   "123",
		},

		{
			Value: int8(126),
			Expected:   "126",
		},
		{
			Value: int8(127),
			Expected:   "127",
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
