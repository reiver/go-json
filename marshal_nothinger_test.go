package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
)

type OptionalString struct {
	value string
	something bool
}
func Nothing(value string) OptionalString {
	return OptionalString{}
}
func Something(value string) OptionalString {
	return OptionalString{
		something: true,
		value: value,
	}
}
func (receiver OptionalString) IsNothing() bool {
	return !(receiver.something)
}
func (receiver OptionalString) MarshalJSON() ([]byte, error) {
	if receiver.IsNothing() {
		return nil, erorr.Error("nothing")
	}

	return json.MarshalString(receiver.value), nil
}

func TestMarshal_nothinger(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: struct {
				Apple   string         `json:"apple,omitempty"`
				Banana  OptionalString `json:"banana,omitempty"`
				Cherry  string         `json:"cherry,"`
			}{},
			Expected: `{"cherry":""}`,
		},



		{
			Value: struct {
				Apple   string         `json:"apple,omitempty"`
				Banana  OptionalString `json:"banana,omitempty"`
				Cherry  string         `json:"cherry,"`
			}{
				Banana: Something(""),
			},
			Expected: `{"banana":"","cherry":""}`,
		},



		{
			Value: struct {
				Apple   string         `json:"apple,omitempty"`
				Banana  OptionalString `json:"banana,omitempty"`
				Cherry  string         `json:"cherry,"`
			}{
				Apple: "first",
				Banana: Something("second"),
				Cherry: "third",
			},
			Expected: `{"apple":"first","banana":"second","cherry":"third"}`,
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
