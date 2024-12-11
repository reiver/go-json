package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_slice(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: []bool(nil),
			Expected: `[]`,
		},
		{
			Value: []int(nil),
			Expected: `[]`,
		},
		{
			Value: []int8(nil),
			Expected: `[]`,
		},
		{
			Value: []int16(nil),
			Expected: `[]`,
		},
		{
			Value: []int32(nil),
			Expected: `[]`,
		},
		{
			Value: []int64(nil),
			Expected: `[]`,
		},
		{
			Value: []string(nil),
			Expected: `[]`,
		},
		{
			Value: []uint(nil),
			Expected: `[]`,
		},
		{
			Value: []uint8(nil),
			Expected: `[]`,
		},
		{
			Value: []uint16(nil),
			Expected: `[]`,
		},
		{
			Value: []uint32(nil),
			Expected: `[]`,
		},
		{
			Value: []uint64(nil),
			Expected: `[]`,
		},









		{
			Value: []bool{},
			Expected: `[]`,
		},
		{
			Value: []int{},
			Expected: `[]`,
		},
		{
			Value: []int8{},
			Expected: `[]`,
		},
		{
			Value: []int16{},
			Expected: `[]`,
		},
		{
			Value: []int32{},
			Expected: `[]`,
		},
		{
			Value: []int64{},
			Expected: `[]`,
		},
		{
			Value: []string{},
			Expected: `[]`,
		},
		{
			Value: []uint{},
			Expected: `[]`,
		},
		{
			Value: []uint8{},
			Expected: `[]`,
		},
		{
			Value: []uint16{},
			Expected: `[]`,
		},
		{
			Value: []uint32{},
			Expected: `[]`,
		},
		{
			Value: []uint64{},
			Expected: `[]`,
		},









		{
			Value: []string{
				"once",
			},
			Expected: `["once"]`,
		},
		{
			Value: []string{
				"once",
				"twice",
			},
			Expected: `["once","twice"]`,
		},
		{
			Value: []string{
				"once",
				"twice",
				"thrice",
			},
			Expected: `["once","twice","thrice"]`,
		},
		{
			Value: []string{
				"once",
				"twice",
				"thrice",
				"fource",
			},
			Expected: `["once","twice","thrice","fource"]`,
		},









		{
			Value: struct {
				First  []string `json:"first"`
				Second []string `json:"second,omitempty"`
				Third  []string `json:"third,omitempty"`
			}{
				First: []string{
					"once",
					"twice",
					"thrice",
					"fource",
				},
				Second: []string{},
			},
			Expected: `{"first":["once","twice","thrice","fource"]}`,
		},
		{
			Value: struct {
				First  []string `json:"first"`
				Second []string `json:"second,omitempty"`
				Third  []string `json:"third,omitempty"`
			}{
				First: []string{
					"once",
					"twice",
					"thrice",
					"fource",
				},
				Second: []string{},
				Third:  []string{"hi"},
			},
			Expected: `{"first":["once","twice","thrice","fource"],"third":["hi"]}`,
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
