package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestUnmarshal_string(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Expected string
	}{
		{
			Bytes:    []byte(`"hello"`),
			Expected: "hello",
		},
		{
			Bytes:    []byte(`""`),
			Expected: "",
		},
		{
			Bytes:    []byte(`"hello\nworld"`),
			Expected: "hello\nworld",
		},
		{
			Bytes:    []byte(`"tab\there"`),
			Expected: "tab\there",
		},
		{
			Bytes:    []byte(`"quote\"inside"`),
			Expected: `quote"inside`,
		},
		{
			Bytes:    []byte(`"🙂"`),
			Expected: "🙂",
		},
	}

	for testNumber, test := range tests {
		var actual string

		err := json.Unmarshal(test.Bytes, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestUnmarshal_bool(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Expected bool
	}{
		{
			Bytes:    []byte(`true`),
			Expected: true,
		},
		{
			Bytes:    []byte(`false`),
			Expected: false,
		},
	}

	for testNumber, test := range tests {
		var actual bool

		err := json.Unmarshal(test.Bytes, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.Expected)
			t.Logf("ACTUAL:   %v", actual)
			continue
		}
	}
}

func TestUnmarshal_int(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes:    []byte(`42`),
			Dst:      func() any { return new(int) },
			Expected: 42,
		},
		{
			Bytes:    []byte(`-100`),
			Dst:      func() any { return new(int) },
			Expected: -100,
		},
		{
			Bytes:    []byte(`0`),
			Dst:      func() any { return new(int) },
			Expected: 0,
		},
		{
			Bytes:    []byte(`127`),
			Dst:      func() any { return new(int8) },
			Expected: int8(127),
		},
		{
			Bytes:    []byte(`-128`),
			Dst:      func() any { return new(int8) },
			Expected: int8(-128),
		},
		{
			Bytes:    []byte(`32767`),
			Dst:      func() any { return new(int16) },
			Expected: int16(32767),
		},
		{
			Bytes:    []byte(`2147483647`),
			Dst:      func() any { return new(int32) },
			Expected: int32(2147483647),
		},
		{
			Bytes:    []byte(`9223372036854775807`),
			Dst:      func() any { return new(int64) },
			Expected: int64(9223372036854775807),
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
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v (%T)", test.Expected, test.Expected)
			t.Logf("ACTUAL:   %#v (%T)", actual, actual)
			continue
		}
	}
}

func TestUnmarshal_uint(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes:    []byte(`42`),
			Dst:      func() any { return new(uint) },
			Expected: uint(42),
		},
		{
			Bytes:    []byte(`255`),
			Dst:      func() any { return new(uint8) },
			Expected: uint8(255),
		},
		{
			Bytes:    []byte(`65535`),
			Dst:      func() any { return new(uint16) },
			Expected: uint16(65535),
		},
		{
			Bytes:    []byte(`4294967295`),
			Dst:      func() any { return new(uint32) },
			Expected: uint32(4294967295),
		},
		{
			Bytes:    []byte(`18446744073709551615`),
			Dst:      func() any { return new(uint64) },
			Expected: uint64(18446744073709551615),
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
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v (%T)", test.Expected, test.Expected)
			t.Logf("ACTUAL:   %#v (%T)", actual, actual)
			continue
		}
	}
}

func TestUnmarshal_float(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Dst      func() any
		Expected any
	}{
		{
			Bytes:    []byte(`3.14`),
			Dst:      func() any { return new(float64) },
			Expected: float64(3.14),
		},
		{
			Bytes:    []byte(`-0.5`),
			Dst:      func() any { return new(float64) },
			Expected: float64(-0.5),
		},
		{
			Bytes:    []byte(`1.5`),
			Dst:      func() any { return new(float32) },
			Expected: float32(1.5),
		},
		{
			Bytes:    []byte(`1e10`),
			Dst:      func() any { return new(float64) },
			Expected: float64(1e10),
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
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v (%T)", test.Expected, test.Expected)
			t.Logf("ACTUAL:   %#v (%T)", actual, actual)
			continue
		}
	}
}

func TestUnmarshal_null(t *testing.T) {

	// Test null into a string pointer.
	var strPtr *string
	err := json.Unmarshal([]byte(`null`), &strPtr)
	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: %s", err)
		return
	}
	if nil != strPtr {
		t.Errorf("Expected nil but got %q.", *strPtr)
	}
}

func TestUnmarshal_any(t *testing.T) {

	tests := []struct{
		Bytes    []byte
		Expected any
	}{
		{
			Bytes:    []byte(`"hello"`),
			Expected: "hello",
		},
		{
			Bytes:    []byte(`42`),
			Expected: json.MustParseNumberString("42"),
		},
		{
			Bytes:    []byte(`true`),
			Expected: true,
		},
		{
			Bytes:    []byte(`false`),
			Expected: false,
		},
		{
			Bytes:    []byte(`null`),
			Expected: nil,
		},
		{
			Bytes:    []byte(`{"key":"value"}`),
			Expected: map[string]any{"key": "value"},
		},
		{
			Bytes:    []byte(`[1,2,3]`),
			Expected: []any{json.MustParseNumberString("1"), json.MustParseNumberString("2"), json.MustParseNumberString("3")},
		},
	}

	for testNumber, test := range tests {
		var actual any

		err := json.Unmarshal(test.Bytes, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v (%T)", test.Expected, test.Expected)
			t.Logf("ACTUAL:   %#v (%T)", actual, actual)
			continue
		}
	}
}
