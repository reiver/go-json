package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_int16(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: int16(-32768),
			Expected:   "-32768",
		},
		{
			Value: int16(-32767),
			Expected:   "-32767",
		},
		{
			Value: int16(-32766),
			Expected:   "-32766",
		},

		{
			Value: int16(-30000),
			Expected:   "-30000",
		},

		{
			Value: int16(-20000),
			Expected:   "-20000",
		},

		{
			Value: int16(-12345),
			Expected:   "-12345",
		},

		{
			Value: int16(-10000),
			Expected:   "-10000",
		},

		{
			Value: int16(-1234),
			Expected:   "-1234",
		},

		{
			Value: int16(-1000),
			Expected:   "-1000",
		},

		{
			Value: int16(-300),
			Expected:   "-300",
		},

		{
			Value: int16(-257),
			Expected:   "-257",
		},
		{
			Value: int16(-256),
			Expected:   "-256",
		},
		{
			Value: int16(-255),
			Expected:   "-255",
		},
		{
			Value: int16(-254),
			Expected:   "-254",
		},

		{
			Value: int16(-200),
			Expected:   "-200",
		},

		{
			Value: int16(-129),
			Expected:   "-129",
		},
		{
			Value: int16(-128),
			Expected:   "-128",
		},
		{
			Value: int16(-127),
			Expected:   "-127",
		},
		{
			Value: int16(-126),
			Expected:   "-126",
		},

		{
			Value: int16(-123),
			Expected:   "-123",
		},

		{
			Value: int16(-100),
			Expected:   "-100",
		},

		{
			Value: int16(-12),
			Expected:   "-12",
		},

		{
			Value: int16(-5),
			Expected:   "-5",
		},
		{
			Value: int16(-4),
			Expected:   "-4",
		},
		{
			Value: int16(-3),
			Expected:   "-3",
		},
		{
			Value: int16(-2),
			Expected:   "-2",
		},
		{
			Value: int16(-1),
			Expected:   "-1",
		},
		{
			Value: int16(0),
			Expected:   "0",
		},
		{
			Value: int16(1),
			Expected:   "1",
		},
		{
			Value: int16(2),
			Expected:   "2",
		},
		{
			Value: int16(3),
			Expected:   "3",
		},
		{
			Value: int16(4),
			Expected:   "4",
		},
		{
			Value: int16(5),
			Expected:   "5",
		},

		{
			Value: int16(12),
			Expected:   "12",
		},

		{
			Value: int16(100),
			Expected:   "100",
		},

		{
			Value: int16(123),
			Expected:   "123",
		},

		{
			Value: int16(126),
			Expected:   "126",
		},
		{
			Value: int16(127),
			Expected:   "127",
		},
		{
			Value: int16(128),
			Expected:   "128",
		},
		{
			Value: int16(129),
			Expected:   "129",
		},

		{
			Value: int16(200),
			Expected:   "200",
		},

		{
			Value: int16(254),
			Expected:   "254",
		},
		{
			Value: int16(255),
			Expected:   "255",
		},
		{
			Value: int16(256),
			Expected:   "256",
		},
		{
			Value: int16(257),
			Expected:   "257",
		},

		{
			Value: int16(300),
			Expected:   "300",
		},

		{
			Value: int16(1000),
			Expected:   "1000",
		},

		{
			Value: int16(1234),
			Expected:   "1234",
		},

		{
			Value: int16(10000),
			Expected:   "10000",
		},

		{
			Value: int16(12345),
			Expected:   "12345",
		},

		{
			Value: int16(20000),
			Expected:   "20000",
		},

		{
			Value: int16(30000),
			Expected:   "30000",
		},

		{
			Value: int16(32766),
			Expected:   "32766",
		},
		{
			Value: int16(32767),
			Expected:   "32767",
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
