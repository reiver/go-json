package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

type OmitAlways1 struct {}
var _ json.OmitAlways = OmitAlways1{}
func (receiver OmitAlways1) JSONOmitAlways(){}

func TestMarshal_omitAlways(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: struct{
				Apple string       `json:"apple"`
				Banana OmitAlways1 `json:"banana"`
				Cherry int         `json:"cherry"`
			}{
				Apple: "one",
			},
			Expected: `{"apple":"one","cherry":0}`,
		},
		{
			Value: map[string]any{
				"apple":"one",
				"banana":OmitAlways1{},
				"cherry":5,
			},
			Expected: `{"apple":"one","cherry":5}`,
		},



		{
			Value: struct {
				Once string   `json:"once"`
				Twice struct{
					Apple string       `json:"apple"`
					Banana OmitAlways1 `json:"banana"`
					Cherry int         `json:"cherry"`
				} `json:"twice"`
				Thrice int    `json:"thrice"`
				Fource uint64 `json:"fource"`
			}{
				Once: "first",
				Twice: struct{
					Apple string       `json:"apple"`
					Banana OmitAlways1 `json:"banana"`
					Cherry int         `json:"cherry"`
				}{
					Apple: "one",
					Banana: OmitAlways1{},
					Cherry: 5,
				},
				Thrice: 3,
				Fource: 4,
			},
			Expected: `{"once":"first","twice":{"apple":"one","cherry":5},"thrice":3,"fource":4}`,
		},
		{
			Value: map[string]any{
				"once":"first",
				"twice":map[string]any{
					"apple":"one",
					"banana":OmitAlways1{},
					"cherry":5,
				},
				"thrice":3,
				"fource":uint64(4),
			},
			Expected: `{"fource":4,"once":"first","thrice":3,"twice":{"apple":"one","cherry":5}}`,
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
