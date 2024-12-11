package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_map(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: map[string]bool(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]int(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]int8(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]int16(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]int32(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]int64(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]string(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]uint(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]uint8(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]uint16(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]uint32(nil),

			Expected: `{}`,
		},
		{
			Value: map[string]uint64(nil),

			Expected: `{}`,
		},









		{
			Value: map[string]bool{},

			Expected: `{}`,
		},
		{
			Value: map[string]int{},

			Expected: `{}`,
		},
		{
			Value: map[string]int8{},

			Expected: `{}`,
		},
		{
			Value: map[string]int16{},

			Expected: `{}`,
		},
		{
			Value: map[string]int32{},

			Expected: `{}`,
		},
		{
			Value: map[string]int64{},

			Expected: `{}`,
		},
		{
			Value: map[string]string{},

			Expected: `{}`,
		},
		{
			Value: map[string]uint{},

			Expected: `{}`,
		},
		{
			Value: map[string]uint8{},

			Expected: `{}`,
		},
		{
			Value: map[string]uint16{},

			Expected: `{}`,
		},
		{
			Value: map[string]uint32{},

			Expected: `{}`,
		},
		{
			Value: map[string]uint64{},

			Expected: `{}`,
		},









		{
			Value: map[string]string{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},



		{
			Value: map[string]any{},

			Expected: `{}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},



		{
			Value: map[any]any{},

			Expected: `{}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},








		{
			Value: map[string]any{
				"apple":true,
				"banana":int(-5),
				"cherry":"5",
				"date":uint(5),
			},
			Expected: `{"apple":true,"banana":-5,"cherry":"5","date":5}`,
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
