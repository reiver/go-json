package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnmarshal_slice(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes: []byte(`["apple","banana","cherry"]`),
			Dst: func() any {
				return &[]string{}
			},
			Expected: []string{"apple", "banana", "cherry"},
		},



		{
			Bytes: []byte(`[1,2,3,4,5]`),
			Dst: func() any {
				return &[]int{}
			},
			Expected: []int{1, 2, 3, 4, 5},
		},



		{
			Bytes: []byte(`[]`),
			Dst: func() any {
				return &[]string{}
			},
			Expected: []string{},
		},



		{
			Bytes: []byte(`["hello",42,true,null]`),
			Dst: func() any {
				return &[]any{}
			},
			Expected: []any{"hello", json.MustParseNumberString("42"), true, nil},
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.Unmarshal(test.Bytes, dst)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		actual := reflect.ValueOf(dst).Elem().Interface()

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", test.Expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}
	}
}

type testSlicePerson struct {
	Name string `json:"name"`
}

func TestUnmarshal_slice_structs(t *testing.T) {

	data := []byte(`[{"name":"alice"},{"name":"bob"}]`)

	var actual []testSlicePerson
	err := json.Unmarshal(data, &actual)
	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: %s", err)
		return
	}

	expected := []testSlicePerson{
		{Name: "alice"},
		{Name: "bob"},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("The actual unmarshaled value is not what was expected.")
		t.Logf("EXPECTED:\n%#v", expected)
		t.Logf("ACTUAL:\n%#v", actual)
	}
}
