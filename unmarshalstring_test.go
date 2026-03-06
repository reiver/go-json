package json_test

import (
	"testing"

	"github.com/reiver/go-json"
)

func TestUnmarshalString(t *testing.T) {
	tests := []struct{
		JSON     []byte
		Expected string
	}{
		{
			JSON: []byte(`""`),
			Expected:     "",
		},



		{
			JSON: []byte(`"hello"`),
			Expected:     "hello",
		},
		{
			JSON: []byte(`"Hello, World!"`),
			Expected:     "Hello, World!",
		},



		{
			JSON: []byte(`"hello world"`),
			Expected:     "hello world",
		},



		{
			JSON: []byte(`"line1\nline2"`),
			Expected:     "line1\nline2",
		},
		{
			JSON: []byte(`"col1\tcol2"`),
			Expected:     "col1\tcol2",
		},
		{
			JSON: []byte(`"back\\slash"`),
			Expected:     "back\\slash",
		},
		{
			JSON: []byte(`"quote\"inside"`),
			Expected:     "quote\"inside",
		},
		{
			JSON: []byte(`"a\/b"`),
			Expected:     "a/b",
		},



		{
			JSON: []byte(`"\u0041\u0042\u0043"`),
			Expected:     "ABC",
		},
		{
			JSON: []byte(`"\u00e9"`),
			Expected:     "\u00e9",
		},



		{
			JSON: []byte(`"café"`),
			Expected:     "café",
		},
		{
			JSON: []byte(`"日本語"`),
			Expected:     "日本語",
		},
		{
			JSON: []byte(`"🎉"`),
			Expected:     "🎉",
		},



		{
			// Persian numerals ۰۱۲۳۴۵۶۷۸۹
			JSON: []byte(`"۰۱۲۳۴۵۶۷۸۹"`),
			Expected:     "۰۱۲۳۴۵۶۷۸۹",
		},
		{
			// Persian numeral via unicode escape (۴ = U+06F4).
			JSON: []byte(`"\u06F4\u06F2"`),
			Expected:     "۴۲",
		},
	}

	for testNumber, test := range tests {
		var actual string
		err := json.UnmarshalString(test.JSON, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual JSON-unmarshaled %T is not what was expected.", testNumber, actual)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}
	}
}

func TestUnmarshalString_fail(t *testing.T) {
	tests := []struct{
		JSON []byte
	}{
		{
			JSON: nil,
		},
		{
			JSON: []byte(""),
		},



		{
			// Unquoted string.
			JSON: []byte(`hello`),
		},



		{
			// Missing closing quote.
			JSON: []byte(`"hello`),
		},
		{
			// Missing opening quote.
			JSON: []byte(`hello"`),
		},



		{
			// Not a string — bool.
			JSON: []byte(`true`),
		},
		{
			// Not a string — number.
			JSON: []byte(`42`),
		},
	}

	for testNumber, test := range tests {
		var value string
		err := json.UnmarshalString(test.JSON, &value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", value)
			t.Logf("JSON: (%d)\n%s", len(test.JSON), test.JSON)
			continue
		}

	}
}
