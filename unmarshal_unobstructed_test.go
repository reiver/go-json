package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnobstructedUnmarshal_trailingComma(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			// Trailing comma in object.
			Bytes: []byte(`{"apple":true,"banana":42,}`),
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



		{
			// Trailing comma in array.
			Bytes: []byte(`[1,2,3,]`),
			Dst: func() any {
				return &[]int{}
			},
			Expected: []int{1, 2, 3},
		},



		{
			// Trailing comma in map.
			Bytes: []byte(`{"a":"x","b":"y",}`),
			Dst: func() any {
				return &map[string]string{}
			},
			Expected: map[string]string{
				"a": "x",
				"b": "y",
			},
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.UnobstructedUnmarshal(test.Bytes, dst)
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

func TestUnobstructedUnmarshal_plusSign(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			// Plus sign on integer into int.
			Bytes: []byte(`+2`),
			Dst: func() any {
				return new(int)
			},
			Expected: 2,
		},



		{
			// Plus sign on integer into any.
			Bytes: []byte(`+42`),
			Dst: func() any {
				var v any
				return &v
			},
			Expected: json.MustParseNumberString("42"),
		},



		{
			// Plus sign with scientific notation.
			Bytes: []byte(`+2E-3`),
			Dst: func() any {
				return new(float64)
			},
			Expected: float64(0.002),
		},



		{
			// Plus sign with decimal.
			Bytes: []byte(`+3.14`),
			Dst: func() any {
				return new(float64)
			},
			Expected: float64(3.14),
		},



		{
			// Plus sign in struct field.
			Bytes: []byte(`{"value":+10}`),
			Dst: func() any {
				return &struct {
					Value int `json:"value"`
				}{}
			},
			Expected: struct {
				Value int `json:"value"`
			}{
				Value: 10,
			},
		},



		{
			// Plus sign in array.
			Bytes: []byte(`[+1,+2,+3]`),
			Dst: func() any {
				return &[]int{}
			},
			Expected: []int{1, 2, 3},
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.UnobstructedUnmarshal(test.Bytes, dst)
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

func TestUnobstructedUnmarshal_leadingZeros(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes: []byte(`07`),
			Dst: func() any {
				return new(int)
			},
			Expected: 7,
		},



		{
			Bytes: []byte(`007`),
			Dst: func() any {
				return new(int)
			},
			Expected: 7,
		},



		{
			Bytes: []byte(`0007`),
			Dst: func() any {
				return new(int)
			},
			Expected: 7,
		},



		{
			Bytes: []byte(`0100`),
			Dst: func() any {
				return new(int)
			},
			Expected: 100,
		},



		{
			Bytes: []byte(`00100`),
			Dst: func() any {
				return new(int)
			},
			Expected: 100,
		},



		{
			// Leading zeros in struct field.
			Bytes: []byte(`{"value":0100}`),
			Dst: func() any {
				return &struct {
					Value int `json:"value"`
				}{}
			},
			Expected: struct {
				Value int `json:"value"`
			}{
				Value: 100,
			},
		},



		{
			// Leading zeros into any.
			Bytes: []byte(`007`),
			Dst: func() any {
				var v any
				return &v
			},
			Expected: json.MustParseNumberString("7"),
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.UnobstructedUnmarshal(test.Bytes, dst)
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

func TestUnobstructedUnmarshal_trailingDot(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			// Trailing dot into int.
			Bytes: []byte(`100.`),
			Dst: func() any {
				return new(int)
			},
			Expected: 100,
		},



		{
			// Trailing dot into float64.
			Bytes: []byte(`100.`),
			Dst: func() any {
				return new(float64)
			},
			Expected: float64(100),
		},



		{
			// Normal decimal into int (already works, for completeness).
			Bytes: []byte(`100.0`),
			Dst: func() any {
				return new(int)
			},
			Expected: 100,
		},



		{
			// Trailing dot in struct field.
			Bytes: []byte(`{"value":100.}`),
			Dst: func() any {
				return &struct {
					Value int `json:"value"`
				}{}
			},
			Expected: struct {
				Value int `json:"value"`
			}{
				Value: 100,
			},
		},



		{
			// Trailing dot into any.
			Bytes: []byte(`100.`),
			Dst: func() any {
				var v any
				return &v
			},
			Expected: json.MustParseNumberString("100"),
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.UnobstructedUnmarshal(test.Bytes, dst)
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

func TestUnobstructedUnmarshal_comments(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			// Line comments.
			Bytes: []byte(
				"{\n" +
				"  // This is a comment\n" +
				"  \"name\": \"hello\",\n" +
				"  \"value\": 42\n" +
				"  // Another comment\n" +
				"}",
			),
			Dst: func() any {
				return &struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
				}{}
			},
			Expected: struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			}{
				Name:  "hello",
				Value: 42,
			},
		},



		{
			// Block comments.
			Bytes: []byte(`{/* comment */"name":"hello",/* another */"value":42}`),
			Dst: func() any {
				return &struct {
					Name  string `json:"name"`
					Value int    `json:"value"`
				}{}
			},
			Expected: struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			}{
				Name:  "hello",
				Value: 42,
			},
		},



		{
			// Comments in array.
			Bytes: []byte("[1, // first\n 2, /* second */ 3]"),
			Dst: func() any {
				return &[]int{}
			},
			Expected: []int{1, 2, 3},
		},
	}

	for testNumber, test := range tests {
		dst := test.Dst()

		err := json.UnobstructedUnmarshal(test.Bytes, dst)
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
