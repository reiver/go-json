package json_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"
)

func TestMergeAndMarshal(t *testing.T) {

	tests := []struct{
		Values []any
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Values: []any{
				struct {
					Apple string
					Banana int
					Cherry bool
				}{},
			},
			Expected: []byte(`{"Apple":"","Banana":0,"Cherry":false}`),
		},
		{
			Values: []any{
				struct {
					Apple string
					Banana int
					Cherry bool
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(`{"Apple":"ONE","Banana":2,"Cherry":true}`),
		},
		{
			Values: []any{
				struct {
					Apple  string `json:"apple"`
					Banana int    `json:"banana"`
					Cherry bool   `json:"cherry"`
				}{
					Apple: "ONE",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(`{"apple":"ONE","banana":2,"cherry":true}`),
		},
		{
			Values: []any{
				struct {
					Apple  string `json:"apple,omitempty"`
					Banana int    `json:"banana"`
					Cherry bool   `json:"cherry"`
				}{
					Apple: "",
					Banana: 2,
					Cherry: true,
				},
			},
			Expected: []byte(`{"banana":2,"cherry":true}`),
		},









		{
			Values: []any{
				struct {
					Apple  string `json:"apple,omitempty"`
					Banana int    `json:"banana"`
					Cherry bool   `json:"cherry"`
				}{
					Apple: "",
					Banana: 2,
					Cherry: true,
				},
				struct {
					Once   string `json:"once"`
					Twice  int    `json:"twice,omitempty"`
					Thrice bool   `json:"thrice,omitempty"`
					Fource string `json:"fource,omitempty"`
				}{
					Once: "",
					Twice: 202,
					Thrice: true,
					Fource: "4",
				},
			},
			Expected: []byte(`{"banana":2,"cherry":true,"once":"","twice":202,"thrice":true,"fource":"4"}`),
		},
		{
			Values: []any{
				struct {
					Apple  string `json:"apple,omitempty"`
					Banana int    `json:"banana"`
					Cherry bool   `json:"cherry"`
				}{
					Apple: "",
					Banana: 2,
					Cherry: true,
				},
				struct {
					Once   string `json:"once"`
					Twice  int    `json:"twice,omitempty"`
					Thrice bool   `json:"thrice,omitempty"`
					Fource string `json:"fource,omitempty"`
				}{
					Once: "",
					Twice: 202,
					Thrice: true,
					Fource: "4",
				},
				map[string]string{
					"m-k1":"v-1",
				},
			},
			Expected: []byte(`{"banana":2,"cherry":true,"once":"","twice":202,"thrice":true,"fource":"4","m-k1":"v-1"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.MergeAndMarshal(test.Values...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("LEN(VALUES) = %d", len(test.Values))
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual json-merged-and-marshaled result is not what was expected", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("LEN(VALUES) = %d", len(test.Values))
			continue
		}
	}
}
