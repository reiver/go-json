package json_test

import (
	"errors"
	"testing"

	"github.com/reiver/go-json"
)

func TestObstructedUnmarshal_unknownField(t *testing.T) {

	type MyStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct{
		Bytes          []byte
		ExpectedErrors int
		ExpectedName   string
		ExpectedAge    int
	}{
		{
			// One unknown field.
			Bytes:          []byte(`{"name":"Alice","age":30,"extra":"ignored"}`),
			ExpectedErrors: 1,
			ExpectedName:   "Alice",
			ExpectedAge:    30,
		},
		{
			// Two unknown fields.
			Bytes:          []byte(`{"name":"Bob","unknown1":1,"age":25,"unknown2":2}`),
			ExpectedErrors: 2,
			ExpectedName:   "Bob",
			ExpectedAge:    25,
		},
		{
			// No unknown fields — no error.
			Bytes:          []byte(`{"name":"Charlie","age":35}`),
			ExpectedErrors: 0,
			ExpectedName:   "Charlie",
			ExpectedAge:    35,
		},
	}

	for testNumber, test := range tests {
		var actual MyStruct

		err := json.ObstructedUnmarshal(test.Bytes, &actual)

		if 0 == test.ExpectedErrors {
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: %s", err)
				continue
			}
		} else {
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
				continue
			}

			var ue json.UnmarshalErrors
			if !errors.As(err, &ue) {
				t.Errorf("For test #%d, expected UnmarshalErrors but got %T.", testNumber, err)
				t.Logf("ERROR: %s", err)
				continue
			}

			if test.ExpectedErrors != len(ue.Errors) {
				t.Errorf("For test #%d, expected %d errors but got %d.", testNumber, test.ExpectedErrors, len(ue.Errors))
				t.Logf("ERRORS: %v", ue.Errors)
				continue
			}

			// Check that each error is an UnknownFieldError.
			for i, e := range ue.Errors {
				var ufe json.UnknownFieldError
				if !errors.As(e, &ufe) {
					t.Errorf("For test #%d, error #%d: expected UnknownFieldError but got %T.", testNumber, i, e)
				}
			}
		}

		if test.ExpectedName != actual.Name {
			t.Errorf("For test #%d, expected Name to be %q but got %q.", testNumber, test.ExpectedName, actual.Name)
		}
		if test.ExpectedAge != actual.Age {
			t.Errorf("For test #%d, expected Age to be %d but got %d.", testNumber, test.ExpectedAge, actual.Age)
		}
	}
}
