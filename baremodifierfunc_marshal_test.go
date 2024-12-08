package json_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
)

func TestBareModifierFunc_struct(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: struct{
				Once string `json:"once,bare"`
			}{
				Once: "123",
			},
			Expected: []byte(`{"once":123}`),
		},
		{
			Value: struct{
				Once  string `json:"once,bare"`
				Twice string `json:"twice"`
			}{
				Once:  "123",
				Twice: "234",
			},
			Expected: []byte(`{"once":123,"twice":"234"}`),
		},
		{
			Value: struct{
				Once   string `json:"once,bare"`
				Twice  string `json:"twice"`
				Thrice string `json:"thrice,bare"`
			}{
				Once:   "123",
				Twice:  "234",
				Thrice: "345",
			},
			Expected: []byte(`{"once":123,"twice":"234","thrice":345}`),
		},
		{
			Value: struct{
				Once    string `json:"once,bare"`
				Twice   string `json:"twice"`
				Thrice  string `json:"thrice,bare"`
				Fource  string `json:"fource"`
			}{
				Once:   "123",
				Twice:  "234",
				Thrice: "345",
				Fource: "456",
			},
			Expected: []byte(`{"once":123,"twice":"234","thrice":345,"fource":"456"}`),
		},



		{
			Value: struct{
				Once    string `json:"once,bare"`
				Twice   string `json:"twice"`
				Thrice  string `json:"thrice,bare"`
				Fource  string `json:"fource"`

				Bad     string `json:"bad,bare"`
			}{
				Once:   "123",
				Twice:  "234",
				Thrice: "345",
				Fource: "456",

				Bad: "apple banana cherry",
			},
			Expected: []byte(`{"once":123,"twice":"234","thrice":345,"fource":"456","bad":apple banana cherry}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one" , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled value is not what was expected" , testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
