package json_test

import (
	"testing"

	"github.com/reiver/go-json"
)

func TestUnmarshalBool(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected bool
	}{
		{
			JSON: []byte(`true`),
			Expected:     true,
		},
		{
			JSON: []byte(`false`),
			Expected:     false,
		},
	}

	for testNumber, test := range tests {
		var actual bool
		err := json.UnmarshalBool(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalBool_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: nil,
		},
		{
			JSON: []byte(""),
		},



		{
			JSON: []byte(`false,`),
		},
		{
			JSON: []byte(`true,`),
		},
		{
			JSON: []byte(`false]`),
		},
		{
			JSON: []byte(`true]`),
		},
		{
			JSON: []byte(`false}`),
		},
		{
			JSON: []byte(`true}`),
		},



		{
			JSON: []byte(`FALSE`),
		},
		{
			JSON: []byte(`TRUE`),
		},
		{
			JSON: []byte(`False`),
		},
		{
			JSON: []byte(`True`),
		},
		{
			JSON: []byte(`t`),
		},
		{
			JSON: []byte(`f`),
		},



		{
			JSON: []byte(`1`),
		},
		{
			JSON: []byte(`0`),
		},



		{
			JSON: []byte(`"true"`),
		},
		{
			JSON: []byte(`"false"`),
		},
	}

	for testNumber, test := range tests {
		var value bool
		err := json.UnmarshalBool(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}
