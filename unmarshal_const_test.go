package json_test

import (
	"errors"
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnmarshal_const_match(t *testing.T) {

	type MyStruct struct {
		Name string            `json:"name"`
		Type json.Const[string] `json:"type" json.value:"Note"`
	}

	tests := []struct{
		Bytes    []byte
		Expected MyStruct
	}{
		{
			// Const field matches expected value — no error.
			Bytes: []byte(`{"name":"hello","type":"Note"}`),
			Expected: MyStruct{
				Name: "hello",
			},
		},
		{
			// Const field absent — no error.
			Bytes: []byte(`{"name":"hello"}`),
			Expected: MyStruct{
				Name: "hello",
			},
		},
	}

	for testNumber, test := range tests {
		var actual MyStruct

		err := json.Unmarshal(test.Bytes, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", test.Expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}
	}
}

func TestUnmarshal_const_mismatch(t *testing.T) {

	type MyStruct struct {
		Name string            `json:"name"`
		Type json.Const[string] `json:"type" json.value:"Note"`
	}

	tests := []struct{
		Bytes          []byte
		ExpectedName   string
		ExpectedErrors int
	}{
		{
			// Const field does not match — error collected, other fields still populated.
			Bytes:          []byte(`{"name":"hello","type":"Article"}`),
			ExpectedName:   "hello",
			ExpectedErrors: 1,
		},
		{
			// Const field has wrong type (number instead of string) — error collected.
			Bytes:          []byte(`{"name":"world","type":123}`),
			ExpectedName:   "world",
			ExpectedErrors: 1,
		},
	}

	for testNumber, test := range tests {
		var actual MyStruct

		err := json.Unmarshal(test.Bytes, &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		// Check that the error is an UnmarshalErrors.
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

		// Check that the error is a ConstMismatchError.
		var cme json.ConstMismatchError
		if !errors.As(ue.Errors[0], &cme) {
			t.Errorf("For test #%d, expected ConstMismatchError but got %T.", testNumber, ue.Errors[0])
			continue
		}

		// Check that the Name field was still populated.
		if test.ExpectedName != actual.Name {
			t.Errorf("For test #%d, expected Name to be %q but got %q.", testNumber, test.ExpectedName, actual.Name)
			continue
		}
	}
}

func TestUnmarshal_const_multiple_mismatches(t *testing.T) {

	type Manitoban struct {
		GivenName    string             `json:"given-name"`
		HomeCountry  json.Const[string] `json:"home-country"  json.value:"Canada"`
		HomeProvince json.Const[string] `json:"home-province" json.value:"Manitoba"`
	}

	// Both const fields mismatch — should collect 2 errors.
	data := []byte(`{"given-name":"Alice","home-country":"USA","home-province":"Texas"}`)

	var actual Manitoban

	err := json.Unmarshal(data, &actual)
	if nil == err {
		t.Errorf("Expected an error but did not get one.")
		return
	}

	var ue json.UnmarshalErrors
	if !errors.As(err, &ue) {
		t.Errorf("Expected UnmarshalErrors but got %T.", err)
		t.Logf("ERROR: %s", err)
		return
	}

	if 2 != len(ue.Errors) {
		t.Errorf("Expected 2 errors but got %d.", len(ue.Errors))
		t.Logf("ERRORS: %v", ue.Errors)
		return
	}

	// Name should still be populated.
	if "Alice" != actual.GivenName {
		t.Errorf("Expected GivenName to be %q but got %q.", "Alice", actual.GivenName)
	}
}

func TestUnmarshal_const_int(t *testing.T) {

	type MyStruct struct {
		Name  string          `json:"name"`
		Value json.Const[int] `json:"value" json.value:"42"`
	}

	tests := []struct{
		Bytes       []byte
		ExpectError bool
	}{
		{
			// Match.
			Bytes:       []byte(`{"name":"hello","value":42}`),
			ExpectError: false,
		},
		{
			// Mismatch.
			Bytes:       []byte(`{"name":"hello","value":99}`),
			ExpectError: true,
		},
	}

	for testNumber, test := range tests {
		var actual MyStruct

		err := json.Unmarshal(test.Bytes, &actual)

		if test.ExpectError {
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
				continue
			}
		} else {
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: %s", err)
				continue
			}
		}

		if "hello" != actual.Name {
			t.Errorf("For test #%d, expected Name to be %q but got %q.", testNumber, "hello", actual.Name)
		}
	}
}
