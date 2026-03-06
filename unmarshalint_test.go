package json_test

import (
	"math"
	"testing"

	"github.com/reiver/go-json"
)

func TestUnmarshalInt(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected int
	}{
		{
			JSON: []byte(`0`),
			Expected:     0,
		},
		{
			JSON: []byte(`1`),
			Expected:     1,
		},
		{
			JSON: []byte(`-1`),
			Expected:     -1,
		},
		{
			JSON: []byte(`42`),
			Expected:     42,
		},
		{
			JSON: []byte(`-42`),
			Expected:     -42,
		},
		{
			JSON: []byte(`12345`),
			Expected:     12345,
		},
		{
			JSON: []byte(`-12345`),
			Expected:     -12345,
		},
	}

	for testNumber, test := range tests {
		var actual int
		err := json.UnmarshalInt(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalInt_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},



		{
			// Floating point.
			JSON: []byte(`1.5`),
		},



		{
			// Not a number — bool.
			JSON: []byte(`true`),
		},
		{
			// Not a number — string.
			JSON: []byte(`"42"`),
		},
		{
			// Not a number — null.
			JSON: []byte(`null`),
		},
	}

	for testNumber, test := range tests {
		var value int
		err := json.UnmarshalInt(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalInt8(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected int8
	}{
		{
			JSON: []byte(`0`),
			Expected:     0,
		},
		{
			JSON: []byte(`1`),
			Expected:     1,
		},
		{
			JSON: []byte(`-1`),
			Expected:     -1,
		},
		{
			JSON: []byte(`127`),
			Expected:     math.MaxInt8,
		},
		{
			JSON: []byte(`-128`),
			Expected:     math.MinInt8,
		},
	}

	for testNumber, test := range tests {
		var actual int8
		err := json.UnmarshalInt8(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalInt8_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Overflow.
			JSON: []byte(`128`),
		},
		{
			// Underflow.
			JSON: []byte(`-129`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value int8
		err := json.UnmarshalInt8(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalInt16(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected int16
	}{
		{
			JSON: []byte(`0`),
			Expected:     0,
		},
		{
			JSON: []byte(`1`),
			Expected:     1,
		},
		{
			JSON: []byte(`-1`),
			Expected:     -1,
		},
		{
			JSON: []byte(`32767`),
			Expected:     math.MaxInt16,
		},
		{
			JSON: []byte(`-32768`),
			Expected:     math.MinInt16,
		},
	}

	for testNumber, test := range tests {
		var actual int16
		err := json.UnmarshalInt16(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalInt16_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Overflow.
			JSON: []byte(`32768`),
		},
		{
			// Underflow.
			JSON: []byte(`-32769`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value int16
		err := json.UnmarshalInt16(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalInt32(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected int32
	}{
		{
			JSON: []byte(`0`),
			Expected:     0,
		},
		{
			JSON: []byte(`1`),
			Expected:     1,
		},
		{
			JSON: []byte(`-1`),
			Expected:     -1,
		},
		{
			JSON: []byte(`2147483647`),
			Expected:     math.MaxInt32,
		},
		{
			JSON: []byte(`-2147483648`),
			Expected:     math.MinInt32,
		},
	}

	for testNumber, test := range tests {
		var actual int32
		err := json.UnmarshalInt32(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalInt32_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Overflow.
			JSON: []byte(`2147483648`),
		},
		{
			// Underflow.
			JSON: []byte(`-2147483649`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value int32
		err := json.UnmarshalInt32(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalInt64(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected int64
	}{
		{
			JSON: []byte(`0`),
			Expected:     0,
		},
		{
			JSON: []byte(`1`),
			Expected:     1,
		},
		{
			JSON: []byte(`-1`),
			Expected:     -1,
		},
		{
			JSON: []byte(`9223372036854775807`),
			Expected:     math.MaxInt64,
		},
		{
			JSON: []byte(`-9223372036854775808`),
			Expected:     math.MinInt64,
		},
	}

	for testNumber, test := range tests {
		var actual int64
		err := json.UnmarshalInt64(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalInt64_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Overflow.
			JSON: []byte(`9223372036854775808`),
		},
		{
			// Underflow.
			JSON: []byte(`-9223372036854775809`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value int64
		err := json.UnmarshalInt64(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}
