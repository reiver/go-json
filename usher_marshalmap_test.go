package json_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
)

func TestUsher_marshalMap(t *testing.T) {

	tests := []struct{
		Map any
		Expected []byte
	}{
		{
			Map: map[string]any(nil),
			Expected: []byte("{}"),
		},
		{
			Map: map[string]int(nil),
			Expected: []byte("{}"),
		},
		{
			Map: map[string]string(nil),
			Expected: []byte("{}"),
		},



		{
			Map: map[string]any{},
			Expected: []byte("{}"),
		},
		{
			Map: map[string]int{},
			Expected: []byte("{}"),
		},
		{
			Map: map[string]string{},
			Expected: []byte("{}"),
		},



		{
			Map: map[string]any{
				"something":true,
			},
			Expected: []byte(`{"something":true}`),
		},
		{
			Map: map[string]int{
				"something":5,
			},
			Expected: []byte(`{"something":5}`),
		},
		{
			Map: map[string]string{
				"something":"five",
			},
			Expected: []byte(`{"something":"five"}`),
		},



		{
			Map: map[string]any{
				"once":1,
				"twice":"two",
				"thrice":3,
				"fource":"four",
			},
			Expected: []byte(`{"fource":"four","once":1,"thrice":3,"twice":"two"}`),
		},
		{
			Map: map[string]int{
				"once":1,
				"twice":2,
				"thrice":3,
				"fource":4,
			},
			Expected: []byte(`{"fource":4,"once":1,"thrice":3,"twice":2}`),
		},
		{
			Map: map[string]string{
				"once":"one",
				"twice":"two",
				"thrice":"three",
				"fource":"four",
			},
			Expected: []byte(`{"fource":"four","once":"one","thrice":"three","twice":"two"}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher

		actual, err := usher.Marshal(test.Map)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("MAP: (%T) %#v", test.Map, test.Map)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled map is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("MAP: (%T) %#v", test.Map, test.Map)
			continue
		}
	}
}
