package json_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-json"
)

type DrinkBottle struct {
	Amount uint // the amount of the drink in it
	Kind string // what type of drink in it
}
func (receiver DrinkBottle) IsEmpty() bool {
	return 0 == receiver.Amount
}
func (receiver DrinkBottle) MarshalJSON() ([]byte, error) {
	return []byte(json.MarshalString(receiver.String())), nil
}
func (receiver DrinkBottle) String() string {
	switch {
	case 0 == receiver.Amount && "" == receiver.Kind:
		return "[drink-bottle] 0 floz"
	case "" == receiver.Kind:
		return fmt.Sprintf("[drink-bottle] %d floz of ?", receiver.Amount)
	default:
		return fmt.Sprintf("[drink-bottle] %d floz of %q", receiver.Amount, receiver.Kind)
	}
}

func TestMarshal_emptier(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value: struct {
				Once   string      `json:"once"`
				Twice  DrinkBottle `json:"twice,omitempty"`
				Thrice DrinkBottle `json:"thrice"`
				Fource string      `json:"fource,omitempty"`
			}{},
			Expected: `{"once":"","thrice":"[drink-bottle] 0 floz"}`,
		},



		{
			Value: struct {
				Once   string      `json:"once"`
				Twice  DrinkBottle `json:"twice,omitempty"`
				Thrice DrinkBottle `json:"thrice"`
				Fource string      `json:"fource,omitempty"`
			}{
				Once: "first",
				Twice: DrinkBottle{Kind:"root beer"},
				Thrice: DrinkBottle{Kind:"cream soda"},
				Fource: "forth",
			},
			Expected: `{"once":"first","thrice":"[drink-bottle] 0 floz of \"cream soda\"","fource":"forth"}`,
		},



		{
			Value: struct {
				Once   string      `json:"once"`
				Twice  DrinkBottle `json:"twice,omitempty"`
				Thrice DrinkBottle `json:"thrice"`
				Fource string      `json:"fource,omitempty"`
			}{
				Once: "first",
				Twice: DrinkBottle{Amount:33},
				Thrice: DrinkBottle{Amount:17},
				Fource: "forth",
			},
			Expected: `{"once":"first","twice":"[drink-bottle] 33 floz of ?","thrice":"[drink-bottle] 17 floz of ?","fource":"forth"}`,
		},



		{
			Value: struct {
				Once   string      `json:"once"`
				Twice  DrinkBottle `json:"twice,omitempty"`
				Thrice DrinkBottle `json:"thrice"`
				Fource string      `json:"fource,omitempty"`
			}{
				Once: "first",
				Twice: DrinkBottle{Kind:"root beer", Amount: 33},
				Thrice: DrinkBottle{Kind:"cream soda", Amount: 17},
				Fource: "forth",
			},
			Expected: `{"once":"first","twice":"[drink-bottle] 33 floz of \"root beer\"","thrice":"[drink-bottle] 17 floz of \"cream soda\"","fource":"forth"}`,
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
