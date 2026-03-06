package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestNumber_String(t *testing.T) {

	tests := []struct{
		Number   json.Number
		Expected string
	}{
		{
			Number:   json.Number{},
			Expected: "0",
		},
		{
			Number:   json.MustParseNumberString("42"),
			Expected: "42",
		},
		{
			Number:   json.MustParseNumberString("-1"),
			Expected: "-1",
		},
		{
			Number:   json.MustParseNumberString("3.14"),
			Expected: "3.14",
		},
		{
			Number:   json.MustParseNumberString("100"),
			Expected: "100",
		},
	}

	for testNumber, test := range tests {
		actual := test.Number.String()

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestNumber_MarshalJSON(t *testing.T) {

	tests := []struct{
		Number   json.Number
		Expected string
	}{
		{
			Number:   json.Number{},
			Expected: "0",
		},
		{
			Number:   json.MustParseNumberString("42"),
			Expected: "42",
		},
		{
			Number:   json.MustParseNumberString("-1"),
			Expected: "-1",
		},
		{
			Number:   json.MustParseNumberString("3.14"),
			Expected: "3.14",
		},
		{
			Number:   json.MustParseNumberString("100"),
			Expected: "100",
		},
	}

	for testNumber, test := range tests {
		actual, err := test.Number.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		if test.Expected != string(actual) {
			t.Errorf("For test #%d, the actual marshaled JSON is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", string(actual))
			continue
		}
	}
}

func TestNumber_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON           string
		ExpectedString string
	}{
		{
			JSON:           "0",
			ExpectedString: "0",
		},
		{
			JSON:           "42",
			ExpectedString: "42",
		},
		{
			JSON:           "-1",
			ExpectedString: "-1",
		},
		{
			JSON:           "3.14",
			ExpectedString: "3.14",
		},
		{
			JSON:           "3.140",
			ExpectedString: "3.14",
		},
		{
			JSON:           "007",
			ExpectedString: "7",
		},



		{
			JSON:          "0100",
			ExpectedString: "100",
		},
		{
			JSON:         "00100",
			ExpectedString: "100",
		},
		{
			JSON:        "000100",
			ExpectedString: "100",
		},
		{
			JSON:       "0000100",
			ExpectedString: "100",
		},

		{
			JSON:           "100.",
			ExpectedString: "100",
		},
		{
			JSON:           "100",
			ExpectedString: "100",
		},
		{
			JSON:           "100.0",
			ExpectedString: "100",
		},
		{
			JSON:           "100.00",
			ExpectedString: "100",
		},
		{
			JSON:           "100.000",
			ExpectedString: "100",
		},
		{
			JSON:           "100.0000",
			ExpectedString: "100",
		},
		{
			JSON:           "3.140",
			ExpectedString: "3.14",
		},
		{
			JSON:                 "3.7e-5",
			ExpectedString: "0.000037",
		},
		{
			JSON:                 "3.7E-5",
			ExpectedString: "0.000037",
		},
		{
			JSON:           "2.8281e7",
			ExpectedString: "28281000",
		},
		{
			JSON:           "2.8281E7",
			ExpectedString: "28281000",
		},
	}

	for testNumber, test := range tests {
		var actual json.Number

		err := actual.UnmarshalJSON([]byte(test.JSON))
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		if test.ExpectedString != actual.String() {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.ExpectedString)
			t.Logf("ACTUAL:   %q", actual.String())
			continue
		}
	}
}

func TestNumber_Unmarshal_any(t *testing.T) {

	tests := []struct{
		Bytes          []byte
		ExpectedString string
	}{
		{
			Bytes:          []byte(`42`),
			ExpectedString: "42",
		},
		{
			Bytes:          []byte(`3.14`),
			ExpectedString: "3.14",
		},
		{
			Bytes:          []byte(`0`),
			ExpectedString: "0",
		},
		{
			Bytes:          []byte(`-100`),
			ExpectedString: "-100",
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

		num, ok := actual.(json.Number)
		if !ok {
			t.Errorf("For test #%d, expected json.Number but got %T.", testNumber, actual)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}

		if test.ExpectedString != num.String() {
			t.Errorf("For test #%d, the actual string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.ExpectedString)
			t.Logf("ACTUAL:   %q", num.String())
			continue
		}
	}
}

func TestNumber_roundtrip(t *testing.T) {

	tests := []struct{
		JSON     string
		Expected string
	}{
		{
			JSON:     `{"value":42}`,
			Expected: `{"value":42}`,
		},
		{
			JSON:     `{"value":3.14}`,
			Expected: `{"value":3.14}`,
		},
	}

	for testNumber, test := range tests {
		var m map[string]any

		err := json.Unmarshal([]byte(test.JSON), &m)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		actual, err := json.Marshal(m)
		if nil != err {
			t.Errorf("For test #%d, did not expect a marshal error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		if test.Expected != string(actual) {
			t.Errorf("For test #%d, the actual JSON is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", test.Expected)
			t.Logf("ACTUAL:   %s", string(actual))
			continue
		}
	}
}

func TestNumber_Int64(t *testing.T) {

	tests := []struct{
		Number      json.Number
		ExpectedInt int64
		ExpectedOK  bool
	}{
		{
			Number:      json.Number{},
			ExpectedInt: 0,
			ExpectedOK:  true,
		},
		{
			Number:      json.MustParseNumberString("42"),
			ExpectedInt: 42,
			ExpectedOK:  true,
		},
		{
			Number:      json.MustParseNumberString("-100"),
			ExpectedInt: -100,
			ExpectedOK:  true,
		},
		{
			Number:      json.MustParseNumberString("9223372036854775807"),
			ExpectedInt: 9223372036854775807,
			ExpectedOK:  true,
		},
		{
			Number:      json.MustParseNumberString("-9223372036854775808"),
			ExpectedInt: -9223372036854775808,
			ExpectedOK:  true,
		},
		{
			Number:      json.MustParseNumberString("3.14"),
			ExpectedInt: 0,
			ExpectedOK:  false,
		},
		{
			Number:      json.MustParseNumberString("18446744073709551615"),
			ExpectedInt: 0,
			ExpectedOK:  false,
		},
	}

	for testNumber, test := range tests {
		actual, ok := test.Number.Int64()

		if test.ExpectedOK != ok {
			t.Errorf("For test #%d, the actual ok is not what was expected.", testNumber)
			t.Logf("EXPECTED OK: %v", test.ExpectedOK)
			t.Logf("ACTUAL OK:   %v", ok)
			continue
		}

		if test.ExpectedInt != actual {
			t.Errorf("For test #%d, the actual int64 is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", test.ExpectedInt)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}

func TestNumber_Uint64(t *testing.T) {

	tests := []struct{
		Number       json.Number
		ExpectedUint uint64
		ExpectedOK   bool
	}{
		{
			Number:       json.Number{},
			ExpectedUint: 0,
			ExpectedOK:   true,
		},
		{
			Number:       json.MustParseNumberString("42"),
			ExpectedUint: 42,
			ExpectedOK:   true,
		},
		{
			Number:       json.MustParseNumberString("18446744073709551615"),
			ExpectedUint: 18446744073709551615,
			ExpectedOK:   true,
		},
		{
			Number:       json.MustParseNumberString("-1"),
			ExpectedUint: 0,
			ExpectedOK:   false,
		},
		{
			Number:       json.MustParseNumberString("3.14"),
			ExpectedUint: 0,
			ExpectedOK:   false,
		},
	}

	for testNumber, test := range tests {
		actual, ok := test.Number.Uint64()

		if test.ExpectedOK != ok {
			t.Errorf("For test #%d, the actual ok is not what was expected.", testNumber)
			t.Logf("EXPECTED OK: %v", test.ExpectedOK)
			t.Logf("ACTUAL OK:   %v", ok)
			continue
		}

		if test.ExpectedUint != actual {
			t.Errorf("For test #%d, the actual uint64 is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", test.ExpectedUint)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}

func TestNumber_Float64(t *testing.T) {

	tests := []struct{
		Number        json.Number
		ExpectedFloat float64
		ExpectedOK    bool
	}{
		{
			Number:        json.Number{},
			ExpectedFloat: 0,
			ExpectedOK:    true,
		},
		{
			Number:        json.MustParseNumberString("42"),
			ExpectedFloat: 42,
			ExpectedOK:    true,
		},
		{
			Number:        json.MustParseNumberString("3.14"),
			ExpectedFloat: 3.14,
			ExpectedOK:    true,
		},
		{
			Number:        json.MustParseNumberString("-0.5"),
			ExpectedFloat: -0.5,
			ExpectedOK:    true,
		},
		{
			Number:        json.MustParseNumberString("100"),
			ExpectedFloat: 100,
			ExpectedOK:    true,
		},
	}

	for testNumber, test := range tests {
		actual, ok := test.Number.Float64()

		if test.ExpectedOK != ok {
			t.Errorf("For test #%d, the actual ok is not what was expected.", testNumber)
			t.Logf("EXPECTED OK: %v", test.ExpectedOK)
			t.Logf("ACTUAL OK:   %v", ok)
			continue
		}

		if test.ExpectedFloat != actual {
			t.Errorf("For test #%d, the actual float64 is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.ExpectedFloat)
			t.Logf("ACTUAL:   %v", actual)
			continue
		}
	}
}

func TestNumber_DeepEqual(t *testing.T) {

	a := json.MustParseNumberString("42")
	b := json.MustParseNumberString("42")

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Two Numbers with same value should be DeepEqual.")
		t.Logf("A: %#v", a)
		t.Logf("B: %#v", b)
	}
}
