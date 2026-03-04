package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_float32(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: float32(0),
			Expected: "0",
		},
		{
			Value: float32(1),
			Expected: "1",
		},
		{
			Value: float32(-1),
			Expected: "-1",
		},
		{
			Value: float32(3.14),
			Expected: "3.140000104904174805",
		},
		{
			Value: float32(-3.14),
			Expected: "-3.140000104904174805",
		},
		{
			Value: float32(0.5),
			Expected: "0.5",
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

func TestMarshal_const_float32(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: struct{
				Something json.Const[float32] `json:"something" json.value:"3.14"`
			}{},
			Expected: `{"something":3.140000104904174805}`,
		},
		{
			Value: struct{
				Something json.Const[float32] `json:"something" json.value:"-2.5"`
			}{},
			Expected: `{"something":-2.5}`,
		},
		{
			Value: struct{
				Something json.Const[float32] `json:"something" json.value:"-25"`
			}{},
			Expected: `{"something":-25}`,
		},
		{
			Value: struct{
				Something json.Const[float32] `json:"something" json.value:"0"`
			}{},
			Expected: `{"something":0}`,
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
