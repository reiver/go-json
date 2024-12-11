package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_uint32(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: uint32(0),
			Expected:    "0",
		},
		{
			Value: uint32(1),
			Expected:    "1",
		},
		{
			Value: uint32(2),
			Expected:    "2",
		},
		{
			Value: uint32(3),
			Expected:    "3",
		},
		{
			Value: uint32(4),
			Expected:    "4",
		},
		{
			Value: uint32(5),
			Expected:    "5",
		},

		{
			Value: uint32(12),
			Expected:    "12",
		},

		{
			Value: uint32(100),
			Expected:    "100",
		},

		{
			Value: uint32(123),
			Expected:    "123",
		},

		{
			Value: uint32(126),
			Expected:    "126",
		},
		{
			Value: uint32(127),
			Expected:    "127",
		},
		{
			Value: uint32(128),
			Expected:    "128",
		},
		{
			Value: uint32(129),
			Expected:    "129",
		},

		{
			Value: uint32(200),
			Expected:    "200",
		},

		{
			Value: uint32(254),
			Expected:    "254",
		},
		{
			Value: uint32(255),
			Expected:    "255",
		},
		{
			Value: uint32(256),
			Expected:    "256",
		},
		{
			Value: uint32(257),
			Expected:    "257",
		},

		{
			Value: uint32(300),
			Expected:    "300",
		},

		{
			Value: uint32(1000),
			Expected:    "1000",
		},

		{
			Value: uint32(1234),
			Expected:    "1234",
		},

		{
			Value: uint32(10000),
			Expected:    "10000",
		},

		{
			Value: uint32(12345),
			Expected:    "12345",
		},

		{
			Value: uint32(20000),
			Expected:    "20000",
		},

		{
			Value: uint32(30000),
			Expected:    "30000",
		},

		{
			Value: uint32(32766),
			Expected:    "32766",
		},
		{
			Value: uint32(32767),
			Expected:    "32767",
		},
		{
			Value: uint32(32768),
			Expected:    "32768",
		},
		{
			Value: uint32(32769),
			Expected:    "32769",
		},
		{
			Value: uint32(32770),
			Expected:    "32770",
		},

		{
			Value: uint32(40000),
			Expected:    "40000",
		},

		{
			Value: uint32(50000),
			Expected:    "50000",
		},

		{
			Value: uint32(60000),
			Expected:    "60000",
		},

		{
			Value: uint32(65534),
			Expected:    "65534",
		},
		{
			Value: uint32(65535),
			Expected:    "65535",
		},
		{
			Value: uint32(65536),
			Expected:    "65536",
		},
		{
			Value: uint32(65537),
			Expected:    "65537",
		},
		{
			Value: uint32(65538),
			Expected:    "65538",
		},
		{
			Value: uint32(65539),
			Expected:    "65539",
		},

		{
			Value: uint32(70000),
			Expected:    "70000",
		},

		{
			Value: uint32(80000),
			Expected:    "80000",
		},

		{
			Value: uint32(90000),
			Expected:    "90000",
		},

		{
			Value: uint32(1000000),
			Expected:    "1000000",
		},

		{
			Value: uint32(10000000),
			Expected:    "10000000",
		},

		{
			Value: uint32(100000000),
			Expected:    "100000000",
		},

		{
			Value: uint32(1000000000),
			Expected:    "1000000000",
		},

		{
			Value: uint32(2000000000),
			Expected:    "2000000000",
		},

		{
			Value: uint32(2147483646),
			Expected:    "2147483646",
		},
		{
			Value: uint32(2147483647),
			Expected:    "2147483647",
		},
		{
			Value: uint32(2147483648),
			Expected:    "2147483648",
		},
		{
			Value: uint32(2147483649),
			Expected:    "2147483649",
		},
		{
			Value: uint32(2147483650),
			Expected:    "2147483650",
		},

		{
			Value: uint32(3000000000),
			Expected:    "3000000000",
		},

		{
			Value: uint32(4000000000),
			Expected:    "4000000000",
		},

		{
			Value: uint32(4294967294),
			Expected:    "4294967294",
		},
		{
			Value: uint32(4294967295),
			Expected:    "4294967295",
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
