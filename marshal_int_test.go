package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_int(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: int(-9223372036854775808),
			Expected:   "-9223372036854775808",
		},
		{
			Value: int(-9223372036854775807),
			Expected:   "-9223372036854775807",
		},
		{
			Value: int(-9223372036854775806),
			Expected:   "-9223372036854775806",
		},

		{
			Value: int(-9000000000000000000),
			Expected:   "-9000000000000000000",
		},

		{
			Value: int(-8000000000000000000),
			Expected:   "-8000000000000000000",
		},

		{
			Value: int(-7000000000000000000),
			Expected:   "-7000000000000000000",
		},

		{
			Value: int(-6000000000000000000),
			Expected:   "-6000000000000000000",
		},

		{
			Value: int(-5000000000000000000),
			Expected:   "-5000000000000000000",
		},

		{
			Value: int(-4000000000000000000),
			Expected:   "-4000000000000000000",
		},

		{
			Value: int(-3000000000000000000),
			Expected:   "-3000000000000000000",
		},

		{
			Value: int(-2000000000000000000),
			Expected:   "-2000000000000000000",
		},

		{
			Value: int(-1000000000000000000),
			Expected:   "-1000000000000000000",
		},

		{
			Value: int(-100000000000000000),
			Expected:   "-100000000000000000",
		},

		{
			Value: int(-10000000000000000),
			Expected:   "-10000000000000000",
		},

		{
			Value: int(-1000000000000000),
			Expected:   "-1000000000000000",
		},

		{
			Value: int(-100000000000000),
			Expected:   "-100000000000000",
		},

		{
			Value: int(-10000000000000),
			Expected:   "-10000000000000",
		},

		{
			Value: int(-1000000000000),
			Expected:   "-1000000000000",
		},

		{
			Value: int(-100000000000),
			Expected:   "-100000000000",
		},

		{
			Value: int(-10000000000),
			Expected:   "-10000000000",
		},

		{
			Value: int(-9000000000),
			Expected:   "-9000000000",
		},

		{
			Value: int(-8000000000),
			Expected:   "-8000000000",
		},

		{
			Value: int(-7000000000),
			Expected:   "-7000000000",
		},

		{
			Value: int(-6000000000),
			Expected:   "-6000000000",
		},

		{
			Value: int(-5000000000),
			Expected:   "-5000000000",
		},

		{
			Value: int(-4294967298),
			Expected:   "-4294967298",
		},
		{
			Value: int(-4294967297),
			Expected:   "-4294967297",
		},
		{
			Value: int(-4294967296),
			Expected:   "-4294967296",
		},
		{
			Value: int(-4294967295),
			Expected:   "-4294967295",
		},
		{
			Value: int(-4294967294),
			Expected:   "-4294967294",
		},

		{
			Value: int(-4000000000),
			Expected:   "-4000000000",
		},

		{
			Value: int(-3000000000),
			Expected:   "-3000000000",
		},

		{
			Value: int(-2147483650),
			Expected:   "-2147483650",
		},
		{
			Value: int(-2147483649),
			Expected:   "-2147483649",
		},
		{
			Value: int(-2147483648),
			Expected:   "-2147483648",
		},
		{
			Value: int(-2147483647),
			Expected:   "-2147483647",
		},
		{
			Value: int(-2147483646),
			Expected:   "-2147483646",
		},

		{
			Value: int(-2000000000),
			Expected:   "-2000000000",
		},

		{
			Value: int(-1000000000),
			Expected:   "-1000000000",
		},

		{
			Value: int(-100000000),
			Expected:   "-100000000",
		},

		{
			Value: int(-10000000),
			Expected:   "-10000000",
		},

		{
			Value: int(-1000000),
			Expected:   "-1000000",
		},

		{
			Value: int(-90000),
			Expected:   "-90000",
		},

		{
			Value: int(-80000),
			Expected:   "-80000",
		},

		{
			Value: int(-70000),
			Expected:   "-70000",
		},

		{
			Value: int(-65539),
			Expected:   "-65539",
		},
		{
			Value: int(-65538),
			Expected:   "-65538",
		},
		{
			Value: int(-65537),
			Expected:   "-65537",
		},
		{
			Value: int(-65536),
			Expected:   "-65536",
		},
		{
			Value: int(-65535),
			Expected:   "-65535",
		},
		{
			Value: int(-65534),
			Expected:   "-65534",
		},

		{
			Value: int(-60000),
			Expected:   "-60000",
		},

		{
			Value: int(-50000),
			Expected:   "-50000",
		},

		{
			Value: int(-40000),
			Expected:   "-40000",
		},

		{
			Value: int(-32770),
			Expected:   "-32770",
		},
		{
			Value: int(-32769),
			Expected:   "-32769",
		},
		{
			Value: int(-32768),
			Expected:   "-32768",
		},
		{
			Value: int(-32767),
			Expected:   "-32767",
		},
		{
			Value: int(-32766),
			Expected:   "-32766",
		},

		{
			Value: int(-30000),
			Expected:   "-30000",
		},

		{
			Value: int(-20000),
			Expected:   "-20000",
		},

		{
			Value: int(-12345),
			Expected:   "-12345",
		},

		{
			Value: int(-10000),
			Expected:   "-10000",
		},

		{
			Value: int(-1234),
			Expected:   "-1234",
		},

		{
			Value: int(-1000),
			Expected:   "-1000",
		},

		{
			Value: int(-300),
			Expected:   "-300",
		},

		{
			Value: int(-257),
			Expected:   "-257",
		},
		{
			Value: int(-256),
			Expected:   "-256",
		},
		{
			Value: int(-255),
			Expected:   "-255",
		},
		{
			Value: int(-254),
			Expected:   "-254",
		},

		{
			Value: int(-200),
			Expected:   "-200",
		},

		{
			Value: int(-129),
			Expected:   "-129",
		},
		{
			Value: int(-128),
			Expected:   "-128",
		},
		{
			Value: int(-127),
			Expected:   "-127",
		},
		{
			Value: int(-126),
			Expected:   "-126",
		},

		{
			Value: int(-123),
			Expected:   "-123",
		},

		{
			Value: int(-100),
			Expected:   "-100",
		},

		{
			Value: int(-12),
			Expected:   "-12",
		},

		{
			Value: int(-5),
			Expected:   "-5",
		},
		{
			Value: int(-4),
			Expected:   "-4",
		},
		{
			Value: int(-3),
			Expected:   "-3",
		},
		{
			Value: int(-2),
			Expected:   "-2",
		},
		{
			Value: int(-1),
			Expected:   "-1",
		},
		{
			Value: int(0),
			Expected:   "0",
		},
		{
			Value: int(1),
			Expected:   "1",
		},
		{
			Value: int(2),
			Expected:   "2",
		},
		{
			Value: int(3),
			Expected:   "3",
		},
		{
			Value: int(4),
			Expected:   "4",
		},
		{
			Value: int(5),
			Expected:   "5",
		},

		{
			Value: int(12),
			Expected:   "12",
		},

		{
			Value: int(100),
			Expected:   "100",
		},

		{
			Value: int(123),
			Expected:   "123",
		},

		{
			Value: int(126),
			Expected:   "126",
		},
		{
			Value: int(127),
			Expected:   "127",
		},
		{
			Value: int(128),
			Expected:   "128",
		},
		{
			Value: int(129),
			Expected:   "129",
		},

		{
			Value: int(200),
			Expected:   "200",
		},

		{
			Value: int(254),
			Expected:   "254",
		},
		{
			Value: int(255),
			Expected:   "255",
		},
		{
			Value: int(256),
			Expected:   "256",
		},
		{
			Value: int(257),
			Expected:   "257",
		},

		{
			Value: int(300),
			Expected:   "300",
		},

		{
			Value: int(1000),
			Expected:   "1000",
		},

		{
			Value: int(1234),
			Expected:   "1234",
		},

		{
			Value: int(10000),
			Expected:   "10000",
		},

		{
			Value: int(12345),
			Expected:   "12345",
		},

		{
			Value: int(20000),
			Expected:   "20000",
		},

		{
			Value: int(30000),
			Expected:   "30000",
		},

		{
			Value: int(32766),
			Expected:   "32766",
		},
		{
			Value: int(32767),
			Expected:   "32767",
		},
		{
			Value: int(32768),
			Expected:   "32768",
		},
		{
			Value: int(32769),
			Expected:   "32769",
		},
		{
			Value: int(32770),
			Expected:   "32770",
		},

		{
			Value: int(40000),
			Expected:   "40000",
		},

		{
			Value: int(50000),
			Expected:   "50000",
		},

		{
			Value: int(60000),
			Expected:   "60000",
		},

		{
			Value: int(65534),
			Expected:   "65534",
		},
		{
			Value: int(65535),
			Expected:   "65535",
		},
		{
			Value: int(65536),
			Expected:   "65536",
		},
		{
			Value: int(65537),
			Expected:   "65537",
		},
		{
			Value: int(65538),
			Expected:   "65538",
		},
		{
			Value: int(65539),
			Expected:   "65539",
		},

		{
			Value: int(70000),
			Expected:   "70000",
		},

		{
			Value: int(80000),
			Expected:   "80000",
		},

		{
			Value: int(90000),
			Expected:   "90000",
		},

		{
			Value: int(1000000),
			Expected:   "1000000",
		},

		{
			Value: int(10000000),
			Expected:   "10000000",
		},

		{
			Value: int(100000000),
			Expected:   "100000000",
		},

		{
			Value: int(1000000000),
			Expected:   "1000000000",
		},

		{
			Value: int(2000000000),
			Expected:   "2000000000",
		},

		{
			Value: int(2147483646),
			Expected:   "2147483646",
		},
		{
			Value: int(2147483647),
			Expected:   "2147483647",
		},
		{
			Value: int(2147483648),
			Expected:   "2147483648",
		},
		{
			Value: int(2147483649),
			Expected:   "2147483649",
		},
		{
			Value: int(2147483650),
			Expected:   "2147483650",
		},

		{
			Value: int(3000000000),
			Expected:   "3000000000",
		},

		{
			Value: int(4000000000),
			Expected:   "4000000000",
		},

		{
			Value: int(4294967294),
			Expected:   "4294967294",
		},
		{
			Value: int(4294967295),
			Expected:   "4294967295",
		},
		{
			Value: int(4294967296),
			Expected:   "4294967296",
		},
		{
			Value: int(4294967297),
			Expected:   "4294967297",
		},
		{
			Value: int(4294967298),
			Expected:   "4294967298",
		},

		{
			Value: int(5000000000),
			Expected:   "5000000000",
		},

		{
			Value: int(6000000000),
			Expected:   "6000000000",
		},

		{
			Value: int(7000000000),
			Expected:   "7000000000",
		},

		{
			Value: int(8000000000),
			Expected:   "8000000000",
		},

		{
			Value: int(8000000000),
			Expected:   "8000000000",
		},

		{
			Value: int(9000000000),
			Expected:   "9000000000",
		},

		{
			Value: int(10000000000),
			Expected:   "10000000000",
		},

		{
			Value: int(100000000000),
			Expected:   "100000000000",
		},

		{
			Value: int(1000000000000),
			Expected:   "1000000000000",
		},

		{
			Value: int(10000000000000),
			Expected:   "10000000000000",
		},

		{
			Value: int(100000000000000),
			Expected:   "100000000000000",
		},

		{
			Value: int(1000000000000000),
			Expected:   "1000000000000000",
		},

		{
			Value: int(10000000000000000),
			Expected:   "10000000000000000",
		},

		{
			Value: int(100000000000000000),
			Expected:   "100000000000000000",
		},

		{
			Value: int(1000000000000000000),
			Expected:   "1000000000000000000",
		},

		{
			Value: int(2000000000000000000),
			Expected:   "2000000000000000000",
		},

		{
			Value: int(3000000000000000000),
			Expected:   "3000000000000000000",
		},

		{
			Value: int(4000000000000000000),
			Expected:   "4000000000000000000",
		},

		{
			Value: int(5000000000000000000),
			Expected:   "5000000000000000000",
		},

		{
			Value: int(6000000000000000000),
			Expected:   "6000000000000000000",
		},

		{
			Value: int(7000000000000000000),
			Expected:   "7000000000000000000",
		},

		{
			Value: int(8000000000000000000),
			Expected:   "8000000000000000000",
		},

		{
			Value: int(9000000000000000000),
			Expected:   "9000000000000000000",
		},

		{
			Value: int(9223372036854775806),
			Expected:   "9223372036854775806",
		},
		{
			Value: int(9223372036854775807),
			Expected:   "9223372036854775807",
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
