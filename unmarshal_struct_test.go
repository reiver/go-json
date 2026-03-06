package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnmarshal_struct(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Expected any
		Dst      func() any // returns a pointer to the destination
	}{
		{
			Bytes: []byte(`{}`),
			Dst: func() any {
				return &struct{}{}
			},
			Expected: struct{}{},
		},



		{
			Bytes: []byte(`{"Apple":true,"Banana":-1,"Cherry":"🙂"}`),
			Dst: func() any {
				return &struct {
					Apple  bool
					Banana int
					Cherry string
				}{}
			},
			Expected: struct {
				Apple  bool
				Banana int
				Cherry string
			}{
				Apple:  true,
				Banana: -1,
				Cherry: "🙂",
			},
		},



		{
			Bytes: []byte(`{"apple":true,"banana":-1,"cherry":"🙂"}`),
			Dst: func() any {
				return &struct {
					Apple  bool   `json:"apple"`
					Banana int    `json:"banana"`
					Cherry string `json:"cherry"`
				}{}
			},
			Expected: struct {
				Apple  bool   `json:"apple"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple:  true,
				Banana: -1,
				Cherry: "🙂",
			},
		},



		// Test json:"-" skip
		{
			Bytes: []byte(`{"banana":-1,"cherry":"hello"}`),
			Dst: func() any {
				return &struct {
					Apple  bool   `json:"-"`
					Banana int    `json:"banana"`
					Cherry string `json:"cherry"`
				}{}
			},
			Expected: struct {
				Apple  bool   `json:"-"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple:  false,
				Banana: -1,
				Cherry: "hello",
			},
		},



		// Test unknown fields are ignored in standard mode
		{
			Bytes: []byte(`{"apple":true,"unknown_field":"ignored","banana":42}`),
			Dst: func() any {
				return &struct {
					Apple  bool `json:"apple"`
					Banana int  `json:"banana"`
				}{}
			},
			Expected: struct {
				Apple  bool `json:"apple"`
				Banana int  `json:"banana"`
			}{
				Apple:  true,
				Banana: 42,
			},
		},



		// Test string with escape sequences
		{
			Bytes: []byte(`{"msg":"hello\nworld","tab":"a\tb"}`),
			Dst: func() any {
				return &struct {
					Msg string `json:"msg"`
					Tab string `json:"tab"`
				}{}
			},
			Expected: struct {
				Msg string `json:"msg"`
				Tab string `json:"tab"`
			}{
				Msg: "hello\nworld",
				Tab: "a\tb",
			},
		},



		// Test whitespace in JSON
		{
			Bytes: []byte(`  {  "apple" : true , "banana" : -1  }  `),
			Dst: func() any {
				return &struct {
					Apple  bool `json:"apple"`
					Banana int  `json:"banana"`
				}{}
			},
			Expected: struct {
				Apple  bool `json:"apple"`
				Banana int  `json:"banana"`
			}{
				Apple:  true,
				Banana: -1,
			},
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

type testInner struct {
	Value int `json:"value"`
}

type testOuter struct {
	Name  string    `json:"name"`
	Inner testInner `json:"inner"`
}

func TestUnmarshal_struct_nested(t *testing.T) {

	data := []byte(`{"name":"outer","inner":{"value":123}}`)

	var actual testOuter
	err := json.Unmarshal(data, &actual)
	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: %s", err)
		return
	}

	expected := testOuter{
		Name:  "outer",
		Inner: testInner{Value: 123},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("The actual unmarshaled value is not what was expected.")
		t.Logf("EXPECTED:\n%#v", expected)
		t.Logf("ACTUAL:\n%#v", actual)
	}
}

type testPointerStruct struct {
	Name  *string `json:"name"`
	Value *int    `json:"value"`
}

func TestUnmarshal_struct_pointer_fields(t *testing.T) {

	{
		data := []byte(`{"name":"hello","value":42}`)

		var actual testPointerStruct
		err := json.Unmarshal(data, &actual)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: %s", err)
			return
		}

		if nil == actual.Name || "hello" != *actual.Name {
			t.Errorf("Expected Name to be \"hello\" but got %v", actual.Name)
		}
		if nil == actual.Value || 42 != *actual.Value {
			t.Errorf("Expected Value to be 42 but got %v", actual.Value)
		}
	}

	{
		data := []byte(`{"name":null,"value":42}`)

		var actual testPointerStruct
		err := json.Unmarshal(data, &actual)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: %s", err)
			return
		}

		if nil != actual.Name {
			t.Errorf("Expected Name to be nil but got %q", *actual.Name)
		}
		if nil == actual.Value || 42 != *actual.Value {
			t.Errorf("Expected Value to be 42 but got %v", actual.Value)
		}
	}
}
