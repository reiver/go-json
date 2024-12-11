package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_uint16(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: uint16(0),
			Expected:    "0",
		},
		{
			Value: uint16(1),
			Expected:    "1",
		},
		{
			Value: uint16(2),
			Expected:    "2",
		},
		{
			Value: uint16(3),
			Expected:    "3",
		},
		{
			Value: uint16(4),
			Expected:    "4",
		},
		{
			Value: uint16(5),
			Expected:    "5",
		},

		{
			Value: uint16(12),
			Expected:    "12",
		},

		{
			Value: uint16(100),
			Expected:    "100",
		},

		{
			Value: uint16(123),
			Expected:    "123",
		},

		{
			Value: uint16(126),
			Expected:    "126",
		},
		{
			Value: uint16(127),
			Expected:    "127",
		},
		{
			Value: uint16(128),
			Expected:    "128",
		},
		{
			Value: uint16(129),
			Expected:    "129",
		},

		{
			Value: uint16(200),
			Expected:    "200",
		},

		{
			Value: uint16(254),
			Expected:    "254",
		},
		{
			Value: uint16(255),
			Expected:    "255",
		},
		{
			Value: uint16(256),
			Expected:    "256",
		},
		{
			Value: uint16(257),
			Expected:    "257",
		},

		{
			Value: uint16(300),
			Expected:    "300",
		},

		{
			Value: uint16(1000),
			Expected:    "1000",
		},

		{
			Value: uint16(1234),
			Expected:    "1234",
		},

		{
			Value: uint16(10000),
			Expected:    "10000",
		},

		{
			Value: uint16(12345),
			Expected:    "12345",
		},

		{
			Value: uint16(20000),
			Expected:    "20000",
		},

		{
			Value: uint16(30000),
			Expected:    "30000",
		},

		{
			Value: uint16(32766),
			Expected:    "32766",
		},
		{
			Value: uint16(32767),
			Expected:    "32767",
		},
		{
			Value: uint16(32768),
			Expected:    "32768",
		},
		{
			Value: uint16(32769),
			Expected:    "32769",
		},
		{
			Value: uint16(32770),
			Expected:    "32770",
		},

		{
			Value: uint16(40000),
			Expected:    "40000",
		},

		{
			Value: uint16(50000),
			Expected:    "50000",
		},

		{
			Value: uint16(60000),
			Expected:    "60000",
		},

		{
			Value: uint16(65534),
			Expected:    "65534",
		},
		{
			Value: uint16(65535),
			Expected:    "65535",
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
