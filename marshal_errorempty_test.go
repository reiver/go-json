package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

type AlwaysErrorEmpty struct{Value int}
func (AlwaysErrorEmpty) MarshalJSON() ([]byte, error) {return nil, json.ErrEmpty("test of ErrorEmpty in test using AlwaysErrorEmpty type")}

func TestMarshal_errorEmpty(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: struct {
				Apple  string           `json:"apple,omitempty"`
				Banana AlwaysErrorEmpty `json:"banana,omitempty"`
				Cherry int              `json:"cherry,omitempty"`
			}{
				Apple: "red",
				Banana: AlwaysErrorEmpty{5},
				Cherry: -5,
			},
			Expected: `{"apple":"red","cherry":-5}`,
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
