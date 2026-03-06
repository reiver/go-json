package json_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
)

func TestUsher_Marshal_omitempty(t *testing.T) {

	tests := []struct{
		Source any
		Expected []byte
	}{
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource"`
			}{},
			Expected: []byte(`{"once":"","twice":0,"thrice":false,"fource":{}}`),
		},



		{
			Source: struct{
				Once string              `json:"once,omitempty"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource"`
			}{},
			Expected: []byte(`{"twice":0,"thrice":false,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice,omitempty"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource"`
			}{},
			Expected: []byte(`{"once":"","thrice":false,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice,omitempty"`
				Fource map[string]string `json:"fource"`
			}{},
			Expected: []byte(`{"once":"","twice":0,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource,omitempty"`
			}{},
			Expected: []byte(`{"once":"","twice":0,"thrice":false}`),
		},



		{
			Source: struct{
				Once string              `json:"once,omitempty"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource"`
			}{
				Fource: map[string]string{},
			},
			Expected: []byte(`{"twice":0,"thrice":false,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice,omitempty"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource"`
			}{
				Fource: map[string]string{},
			},
			Expected: []byte(`{"once":"","thrice":false,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice,omitempty"`
				Fource map[string]string `json:"fource"`
			}{
				Fource: map[string]string{},
			},
			Expected: []byte(`{"once":"","twice":0,"fource":{}}`),
		},
		{
			Source: struct{
				Once string              `json:"once"`
				Twice int                `json:"twice"`
				Thrice bool              `json:"thrice"`
				Fource map[string]string `json:"fource,omitempty"`
			}{
				Fource: map[string]string{},
			},
			Expected: []byte(`{"once":"","twice":0,"thrice":false}`),
		},
	}

	for testNumber, test := range tests {

		var usher json.Usher

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-marshaled map is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T)\n%#v", test.Source, test.Source)
			continue
		}
	}
}

// Modifier that returns ErrorEmpty should cause the field to be omitted when omitempty is set.
func TestUsher_Marshal_omitemptyErrorEmpty(t *testing.T) {
	tests := []struct{
		Source any
		Expected []byte
	}{
		{
			Source: struct{
				Name  string `json:"name"`
				Value string `json:"value,omitempty,emptymod"`
			}{
				Name: "hello",
				Value: "world",
			},
			Expected: []byte(`{"name":"hello"}`),
		},
		{
			Source: struct{
				Value string `json:"value,omitempty,emptymod"`
				Name  string `json:"name"`
			}{
				Name: "hello",
				Value: "world",
			},
			Expected: []byte(`{"name":"hello"}`),
		},
		// Without omitempty, the modifier error should propagate.
	}

	for testNumber, test := range tests {

		var usher json.Usher
		usher.ImplantModifier("emptymod", func([]byte) ([]byte, error) {
			return nil, json.ErrEmpty("test empty")
		}, nil)

		actual, err := usher.Marshal(test.Source)
		if nil != err {
			t.Errorf("For mtest #%d, did not expect an error but actually got one.", testNumber)
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
