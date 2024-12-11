package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_uint64(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: uint64(0),
			Expected:    "0",
		},
		{
			Value: uint64(1),
			Expected:    "1",
		},
		{
			Value: uint64(2),
			Expected:    "2",
		},
		{
			Value: uint64(3),
			Expected:    "3",
		},
		{
			Value: uint64(4),
			Expected:    "4",
		},
		{
			Value: uint64(5),
			Expected:    "5",
		},

		{
			Value: uint64(12),
			Expected:    "12",
		},

		{
			Value: uint64(100),
			Expected:    "100",
		},

		{
			Value: uint64(123),
			Expected:    "123",
		},

		{
			Value: uint64(126),
			Expected:    "126",
		},
		{
			Value: uint64(127),
			Expected:    "127",
		},
		{
			Value: uint64(128),
			Expected:    "128",
		},
		{
			Value: uint64(129),
			Expected:    "129",
		},

		{
			Value: uint64(200),
			Expected:    "200",
		},

		{
			Value: uint64(254),
			Expected:    "254",
		},
		{
			Value: uint64(255),
			Expected:    "255",
		},
		{
			Value: uint64(256),
			Expected:    "256",
		},
		{
			Value: uint64(257),
			Expected:    "257",
		},

		{
			Value: uint64(300),
			Expected:    "300",
		},

		{
			Value: uint64(1000),
			Expected:    "1000",
		},

		{
			Value: uint64(1234),
			Expected:    "1234",
		},

		{
			Value: uint64(10000),
			Expected:    "10000",
		},

		{
			Value: uint64(12345),
			Expected:    "12345",
		},

		{
			Value: uint64(20000),
			Expected:    "20000",
		},

		{
			Value: uint64(30000),
			Expected:    "30000",
		},

		{
			Value: uint64(32766),
			Expected:    "32766",
		},
		{
			Value: uint64(32767),
			Expected:    "32767",
		},
		{
			Value: uint64(32768),
			Expected:    "32768",
		},
		{
			Value: uint64(32769),
			Expected:    "32769",
		},
		{
			Value: uint64(32770),
			Expected:    "32770",
		},

		{
			Value: uint64(40000),
			Expected:    "40000",
		},

		{
			Value: uint64(50000),
			Expected:    "50000",
		},

		{
			Value: uint64(60000),
			Expected:    "60000",
		},

		{
			Value: uint64(65534),
			Expected:    "65534",
		},
		{
			Value: uint64(65535),
			Expected:    "65535",
		},
		{
			Value: uint64(65536),
			Expected:    "65536",
		},
		{
			Value: uint64(65537),
			Expected:    "65537",
		},
		{
			Value: uint64(65538),
			Expected:    "65538",
		},
		{
			Value: uint64(65539),
			Expected:    "65539",
		},

		{
			Value: uint64(70000),
			Expected:    "70000",
		},

		{
			Value: uint64(80000),
			Expected:    "80000",
		},

		{
			Value: uint64(90000),
			Expected:    "90000",
		},

		{
			Value: uint64(1000000),
			Expected:    "1000000",
		},

		{
			Value: uint64(10000000),
			Expected:    "10000000",
		},

		{
			Value: uint64(100000000),
			Expected:    "100000000",
		},

		{
			Value: uint64(1000000000),
			Expected:    "1000000000",
		},

		{
			Value: uint64(2000000000),
			Expected:    "2000000000",
		},

		{
			Value: uint64(2147483646),
			Expected:    "2147483646",
		},
		{
			Value: uint64(2147483647),
			Expected:    "2147483647",
		},
		{
			Value: uint64(2147483648),
			Expected:    "2147483648",
		},
		{
			Value: uint64(2147483649),
			Expected:    "2147483649",
		},
		{
			Value: uint64(2147483650),
			Expected:    "2147483650",
		},

		{
			Value: uint64(3000000000),
			Expected:    "3000000000",
		},

		{
			Value: uint64(4000000000),
			Expected:    "4000000000",
		},

		{
			Value: uint64(4294967294),
			Expected:    "4294967294",
		},
		{
			Value: uint64(4294967295),
			Expected:    "4294967295",
		},
		{
			Value: uint64(4294967296),
			Expected:    "4294967296",
		},
		{
			Value: uint64(4294967297),
			Expected:    "4294967297",
		},
		{
			Value: uint64(4294967298),
			Expected:    "4294967298",
		},

		{
			Value: uint64(5000000000),
			Expected:    "5000000000",
		},

		{
			Value: uint64(6000000000),
			Expected:    "6000000000",
		},

		{
			Value: uint64(7000000000),
			Expected:    "7000000000",
		},

		{
			Value: uint64(8000000000),
			Expected:    "8000000000",
		},

		{
			Value: uint64(8000000000),
			Expected:    "8000000000",
		},

		{
			Value: uint64(9000000000),
			Expected:    "9000000000",
		},

		{
			Value: uint64(10000000000),
			Expected:    "10000000000",
		},

		{
			Value: uint64(100000000000),
			Expected:    "100000000000",
		},

		{
			Value: uint64(1000000000000),
			Expected:    "1000000000000",
		},

		{
			Value: uint64(10000000000000),
			Expected:    "10000000000000",
		},

		{
			Value: uint64(100000000000000),
			Expected:    "100000000000000",
		},

		{
			Value: uint64(1000000000000000),
			Expected:    "1000000000000000",
		},

		{
			Value: uint64(10000000000000000),
			Expected:    "10000000000000000",
		},

		{
			Value: uint64(100000000000000000),
			Expected:    "100000000000000000",
		},

		{
			Value: uint64(1000000000000000000),
			Expected:    "1000000000000000000",
		},

		{
			Value: uint64(2000000000000000000),
			Expected:    "2000000000000000000",
		},

		{
			Value: uint64(3000000000000000000),
			Expected:    "3000000000000000000",
		},

		{
			Value: uint64(4000000000000000000),
			Expected:    "4000000000000000000",
		},

		{
			Value: uint64(5000000000000000000),
			Expected:    "5000000000000000000",
		},

		{
			Value: uint64(6000000000000000000),
			Expected:    "6000000000000000000",
		},

		{
			Value: uint64(7000000000000000000),
			Expected:    "7000000000000000000",
		},

		{
			Value: uint64(8000000000000000000),
			Expected:    "8000000000000000000",
		},

		{
			Value: uint64(9000000000000000000),
			Expected:    "9000000000000000000",
		},

		{
			Value: uint64(9223372036854775806),
			Expected:    "9223372036854775806",
		},
		{
			Value: uint64(9223372036854775807),
			Expected:    "9223372036854775807",
		},

		{
			Value: uint64(10000000000000000000),
			Expected :   "10000000000000000000",
		},

		{
			Value: uint64(12345678909876543210),
			Expected :   "12345678909876543210",
		},

		{
			Value: uint64(18446744073709551613),
			Expected :   "18446744073709551613",
		},
		{
			Value: uint64(18446744073709551614),
			Expected :   "18446744073709551614",
		},
		{
			Value: uint64(18446744073709551615),
			Expected :   "18446744073709551615",
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
