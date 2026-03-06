package json_test

import (
	"math"
	"testing"

	"github.com/reiver/go-json"
)

func TestUnmarshalUint(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected uint
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
			JSON: []byte(`42`),
			Expected:     42,
		},
		{
			JSON: []byte(`12345`),
			Expected:     12345,
		},
	}

	for testNumber, test := range tests {
		var actual uint
		err := json.UnmarshalUint(test.JSON, &actual)
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

func TestUnmarshalUint_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},



		{
			// Negative number.
			JSON: []byte(`-1`),
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
		var value uint
		err := json.UnmarshalUint(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalUint8(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected uint8
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
			JSON: []byte(`255`),
			Expected:     math.MaxUint8,
		},
	}

	for testNumber, test := range tests {
		var actual uint8
		err := json.UnmarshalUint8(test.JSON, &actual)
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

func TestUnmarshalUint8_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Negative number.
			JSON: []byte(`-1`),
		},
		{
			// Overflow.
			JSON: []byte(`256`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value uint8
		err := json.UnmarshalUint8(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalUint16(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected uint16
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
			JSON: []byte(`65535`),
			Expected:     math.MaxUint16,
		},
	}

	for testNumber, test := range tests {
		var actual uint16
		err := json.UnmarshalUint16(test.JSON, &actual)
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

func TestUnmarshalUint16_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Negative number.
			JSON: []byte(`-1`),
		},
		{
			// Overflow.
			JSON: []byte(`65536`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value uint16
		err := json.UnmarshalUint16(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalUint32(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected uint32
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
			JSON: []byte(`4294967295`),
			Expected:     math.MaxUint32,
		},
	}

	for testNumber, test := range tests {
		var actual uint32
		err := json.UnmarshalUint32(test.JSON, &actual)
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

func TestUnmarshalUint32_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Negative number.
			JSON: []byte(`-1`),
		},
		{
			// Overflow.
			JSON: []byte(`4294967296`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value uint32
		err := json.UnmarshalUint32(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}

func TestUnmarshalUint64(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected uint64
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
			JSON: []byte(`18446744073709551615`),
			Expected:     math.MaxUint64,
		},
	}

	for testNumber, test := range tests {
		var actual uint64
		err := json.UnmarshalUint64(test.JSON, &actual)
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

func TestUnmarshalUint64_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: []byte(``),
		},
		{
			// Negative number.
			JSON: []byte(`-1`),
		},
		{
			// Overflow.
			JSON: []byte(`18446744073709551616`),
		},
		{
			// Floating point.
			JSON: []byte(`1.5`),
		},
	}

	for testNumber, test := range tests {
		var value uint64
		err := json.UnmarshalUint64(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}
