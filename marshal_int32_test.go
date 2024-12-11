package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_int32(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: int32(-2147483648),
			Expected:   "-2147483648",
		},
		{
			Value: int32(-2147483647),
			Expected:   "-2147483647",
		},
		{
			Value: int32(-2147483646),
			Expected:   "-2147483646",
		},

		{
			Value: int32(-2000000000),
			Expected:   "-2000000000",
		},

		{
			Value: int32(-1000000000),
			Expected:   "-1000000000",
		},

		{
			Value: int32(-100000000),
			Expected:   "-100000000",
		},

		{
			Value: int32(-10000000),
			Expected:   "-10000000",
		},

		{
			Value: int32(-1000000),
			Expected:   "-1000000",
		},

		{
			Value: int32(-90000),
			Expected:   "-90000",
		},

		{
			Value: int32(-80000),
			Expected:   "-80000",
		},

		{
			Value: int32(-70000),
			Expected:   "-70000",
		},

		{
			Value: int32(-65539),
			Expected:   "-65539",
		},
		{
			Value: int32(-65538),
			Expected:   "-65538",
		},
		{
			Value: int32(-65537),
			Expected:   "-65537",
		},
		{
			Value: int32(-65536),
			Expected:   "-65536",
		},
		{
			Value: int32(-65535),
			Expected:   "-65535",
		},
		{
			Value: int32(-65534),
			Expected:   "-65534",
		},

		{
			Value: int32(-60000),
			Expected:   "-60000",
		},

		{
			Value: int32(-50000),
			Expected:   "-50000",
		},

		{
			Value: int32(-40000),
			Expected:   "-40000",
		},

		{
			Value: int32(-32770),
			Expected:   "-32770",
		},
		{
			Value: int32(-32769),
			Expected:   "-32769",
		},
		{
			Value: int32(-32768),
			Expected:   "-32768",
		},
		{
			Value: int32(-32767),
			Expected:   "-32767",
		},
		{
			Value: int32(-32766),
			Expected:   "-32766",
		},

		{
			Value: int32(-30000),
			Expected:   "-30000",
		},

		{
			Value: int32(-20000),
			Expected:   "-20000",
		},

		{
			Value: int32(-12345),
			Expected:   "-12345",
		},

		{
			Value: int32(-10000),
			Expected:   "-10000",
		},

		{
			Value: int32(-1234),
			Expected:   "-1234",
		},

		{
			Value: int32(-1000),
			Expected:   "-1000",
		},

		{
			Value: int32(-300),
			Expected:   "-300",
		},

		{
			Value: int32(-257),
			Expected:   "-257",
		},
		{
			Value: int32(-256),
			Expected:   "-256",
		},
		{
			Value: int32(-255),
			Expected:   "-255",
		},
		{
			Value: int32(-254),
			Expected:   "-254",
		},

		{
			Value: int32(-200),
			Expected:   "-200",
		},

		{
			Value: int32(-129),
			Expected:   "-129",
		},
		{
			Value: int32(-128),
			Expected:   "-128",
		},
		{
			Value: int32(-127),
			Expected:   "-127",
		},
		{
			Value: int32(-126),
			Expected:   "-126",
		},

		{
			Value: int32(-123),
			Expected:   "-123",
		},

		{
			Value: int32(-100),
			Expected:   "-100",
		},

		{
			Value: int32(-12),
			Expected:   "-12",
		},

		{
			Value: int32(-5),
			Expected:   "-5",
		},
		{
			Value: int32(-4),
			Expected:   "-4",
		},
		{
			Value: int32(-3),
			Expected:   "-3",
		},
		{
			Value: int32(-2),
			Expected:   "-2",
		},
		{
			Value: int32(-1),
			Expected:   "-1",
		},
		{
			Value: int32(0),
			Expected:   "0",
		},
		{
			Value: int32(1),
			Expected:   "1",
		},
		{
			Value: int32(2),
			Expected:   "2",
		},
		{
			Value: int32(3),
			Expected:   "3",
		},
		{
			Value: int32(4),
			Expected:   "4",
		},
		{
			Value: int32(5),
			Expected:   "5",
		},

		{
			Value: int32(12),
			Expected:   "12",
		},

		{
			Value: int32(100),
			Expected:   "100",
		},

		{
			Value: int32(123),
			Expected:   "123",
		},

		{
			Value: int32(126),
			Expected:   "126",
		},
		{
			Value: int32(127),
			Expected:   "127",
		},
		{
			Value: int32(128),
			Expected:   "128",
		},
		{
			Value: int32(129),
			Expected:   "129",
		},

		{
			Value: int32(200),
			Expected:   "200",
		},

		{
			Value: int32(254),
			Expected:   "254",
		},
		{
			Value: int32(255),
			Expected:   "255",
		},
		{
			Value: int32(256),
			Expected:   "256",
		},
		{
			Value: int32(257),
			Expected:   "257",
		},

		{
			Value: int32(300),
			Expected:   "300",
		},

		{
			Value: int32(1000),
			Expected:   "1000",
		},

		{
			Value: int32(1234),
			Expected:   "1234",
		},

		{
			Value: int32(10000),
			Expected:   "10000",
		},

		{
			Value: int32(12345),
			Expected:   "12345",
		},

		{
			Value: int32(20000),
			Expected:   "20000",
		},

		{
			Value: int32(30000),
			Expected:   "30000",
		},

		{
			Value: int32(32766),
			Expected:   "32766",
		},
		{
			Value: int32(32767),
			Expected:   "32767",
		},
		{
			Value: int32(32768),
			Expected:   "32768",
		},
		{
			Value: int32(32769),
			Expected:   "32769",
		},
		{
			Value: int32(32770),
			Expected:   "32770",
		},

		{
			Value: int32(40000),
			Expected:   "40000",
		},

		{
			Value: int32(50000),
			Expected:   "50000",
		},

		{
			Value: int32(60000),
			Expected:   "60000",
		},

		{
			Value: int32(65534),
			Expected:   "65534",
		},
		{
			Value: int32(65535),
			Expected:   "65535",
		},
		{
			Value: int32(65536),
			Expected:   "65536",
		},
		{
			Value: int32(65537),
			Expected:   "65537",
		},
		{
			Value: int32(65538),
			Expected:   "65538",
		},
		{
			Value: int32(65539),
			Expected:   "65539",
		},

		{
			Value: int32(70000),
			Expected:   "70000",
		},

		{
			Value: int32(80000),
			Expected:   "80000",
		},

		{
			Value: int32(90000),
			Expected:   "90000",
		},

		{
			Value: int32(1000000),
			Expected:   "1000000",
		},

		{
			Value: int32(10000000),
			Expected:   "10000000",
		},

		{
			Value: int32(100000000),
			Expected:   "100000000",
		},

		{
			Value: int32(1000000000),
			Expected:   "1000000000",
		},

		{
			Value: int32(2000000000),
			Expected:   "2000000000",
		},

		{
			Value: int32(2147483646),
			Expected:   "2147483646",
		},
		{
			Value: int32(2147483647),
			Expected:   "2147483647",
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
