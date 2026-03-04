package json_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
)

func TestUsher_Marshal_nullempty(t *testing.T) {

	tests := []struct{
		Source any
		Expected []byte
	}{
		// Zero-value fields with nullempty should produce null.
		{
			Source: struct{
				Once string              `json:"once,nullempty"`
				Twice int                `json:"twice"`
			}{},
			Expected: []byte(`{"once":null,"twice":0}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice,nullempty"`
			}{},
			Expected: []byte(`{"once":"","twice":null}`),
		},
		{
			Source: struct{
				Once bool `json:"once,nullempty"`
			}{},
			Expected: []byte(`{"once":null}`),
		},
		{
			Source: struct{
				Once map[string]string `json:"once,nullempty"`
			}{
				Once: map[string]string{},
			},
			Expected: []byte(`{"once":null}`),
		},
		{
			Source: struct{
				Once []string `json:"once,nullempty"`
			}{
				Once: []string{},
			},
			Expected: []byte(`{"once":null}`),
		},



		// Non-empty fields with nullempty should marshal normally.
		{
			Source: struct{
				Once string `json:"once,nullempty"`
				Twice int   `json:"twice,nullempty"`
			}{
				Once: "hello",
				Twice: 42,
			},
			Expected: []byte(`{"once":"hello","twice":42}`),
		},
		{
			Source: struct{
				Once bool `json:"once,nullempty"`
			}{
				Once: true,
			},
			Expected: []byte(`{"once":true}`),
		},
		{
			Source: struct{
				Once map[string]string `json:"once,nullempty"`
			}{
				Once: map[string]string{"a": "b"},
			},
			Expected: []byte(`{"once":{"a":"b"}}`),
		},
		{
			Source: struct{
				Once []string `json:"once,nullempty"`
			}{
				Once: []string{"x"},
			},
			Expected: []byte(`{"once":["x"]}`),
		},



		// Multiple fields, some empty some not.
		{
			Source: struct{
				Once   string `json:"once,nullempty"`
				Twice  int    `json:"twice,nullempty"`
				Thrice bool   `json:"thrice"`
			}{
				Twice: 5,
			},
			Expected: []byte(`{"once":null,"twice":5,"thrice":false}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled struct is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}
	}
}

// Test that nullempty with Emptier interface produces null.
type nullemptyEmptier struct {
	empty bool
}
func (n nullemptyEmptier) IsEmpty() bool { return n.empty }
func (n nullemptyEmptier) MarshalJSON() ([]byte, error) { return []byte(`"not-empty"`), nil }

func TestUsher_Marshal_nullempty_emptier(t *testing.T) {

	tests := []struct{
		Source any
		Expected []byte
	}{
		{
			Source: struct{
				Value nullemptyEmptier `json:"value,nullempty"`
			}{
				Value: nullemptyEmptier{empty: true},
			},
			Expected: []byte(`{"value":null}`),
		},
		{
			Source: struct{
				Value nullemptyEmptier `json:"value,nullempty"`
			}{
				Value: nullemptyEmptier{empty: false},
			},
			Expected: []byte(`{"value":"not-empty"}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled struct is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}
	}
}

// Test that nullempty with Nothinger interface produces null.
type nullemptyNothinger struct {
	nothing bool
}
func (n nullemptyNothinger) IsNothing() bool { return n.nothing }
func (n nullemptyNothinger) MarshalJSON() ([]byte, error) { return []byte(`"something"`), nil }

func TestUsher_Marshal_nullempty_nothinger(t *testing.T) {

	tests := []struct{
		Source any
		Expected []byte
	}{
		{
			Source: struct{
				Value nullemptyNothinger `json:"value,nullempty"`
			}{
				Value: nullemptyNothinger{nothing: true},
			},
			Expected: []byte(`{"value":null}`),
		},
		{
			Source: struct{
				Value nullemptyNothinger `json:"value,nullempty"`
			}{
				Value: nullemptyNothinger{nothing: false},
			},
			Expected: []byte(`{"value":"something"}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled struct is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}
	}
}

// Modifier that returns ErrorEmpty with nullempty should produce null instead of omitting.
func TestUsher_Marshal_nullemptyErrorEmpty(t *testing.T) {
	tests := []struct{
		Source any
		Expected []byte
	}{
		{
			Source: struct{
				Name  string `json:"name"`
				Value string `json:"value,nullempty,emptymod"`
			}{
				Name: "hello",
				Value: "world",
			},
			Expected: []byte(`{"name":"hello","value":null}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher
		usher.ImplantModifier("emptymod", func([]byte) ([]byte, error) {
			return nil, json.ErrEmpty("test empty")
		})

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled struct is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}
	}
}
