package json

import (
	"testing"
)

func TestUnquoteString(t *testing.T) {

	tests := []struct{
		Input    []byte
		Expected string
	}{
		// Simple strings (fast path — no escapes).
		{
			Input:    []byte(`"hello"`),
			Expected: "hello",
		},
		{
			Input:    []byte(`""`),
			Expected: "",
		},
		{
			Input:    []byte(`"abc def"`),
			Expected: "abc def",
		},
		{
			Input:    []byte(`"🙂"`),
			Expected: "🙂",
		},



		// Basic escape sequences.
		{
			Input:    []byte(`"hello\nworld"`),
			Expected: "hello\nworld",
		},
		{
			Input:    []byte(`"tab\there"`),
			Expected: "tab\there",
		},
		{
			Input:    []byte(`"quote\"inside"`),
			Expected: `quote"inside`,
		},
		{
			Input:    []byte(`"back\\slash"`),
			Expected: `back\slash`,
		},
		{
			Input:    []byte(`"for\/ward"`),
			Expected: "for/ward",
		},
		{
			Input:    []byte(`"bs\b"`),
			Expected: "bs\b",
		},
		{
			Input:    []byte(`"ff\f"`),
			Expected: "ff\f",
		},
		{
			Input:    []byte(`"cr\r"`),
			Expected: "cr\r",
		},



		// Multiple escapes in one string.
		{
			Input:    []byte(`"a\tb\nc"`),
			Expected: "a\tb\nc",
		},
		{
			Input:    []byte(`"\\\""`),
			Expected: `\"`,
		},



		// Unicode escapes (\uXXXX).
		{
			Input:    []byte(`"\u0041"`),
			Expected: "A",
		},
		{
			Input:    []byte(`"\u00e9"`),
			Expected: "\u00e9", // é
		},
		{
			Input:    []byte(`"\u00E9"`),
			Expected: "\u00e9", // é — uppercase hex
		},
		{
			Input:    []byte(`"\u4e16\u754c"`),
			Expected: "世界",
		},
		{
			Input:    []byte(`"abc\u0020def"`),
			Expected: "abc def",
		},



		// Surrogate pairs for code points above U+FFFF.
		{
			// U+1F600 (😀) = \uD83D\uDE00
			Input:    []byte(`"\uD83D\uDE00"`),
			Expected: "😀",
		},
		{
			// U+1F4A9 (💩) = \uD83D\uDCA9
			Input:    []byte(`"\uD83D\uDCA9"`),
			Expected: "💩",
		},
		{
			// U+10000 (𐀀) = \uD800\uDC00
			Input:    []byte(`"\uD800\uDC00"`),
			Expected: "\U00010000",
		},



		// Mixed content.
		{
			Input:    []byte(`"hello \u0041 world\n\ttab"`),
			Expected: "hello A world\n\ttab",
		},
	}

	for testNumber, test := range tests {
		actual, err := unquoteString(test.Input)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("INPUT: %s", test.Input)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("INPUT:    %s", test.Input)
			continue
		}
	}
}

func TestUnquoteString_fail(t *testing.T) {

	tests := []struct{
		Input []byte
	}{
		// Missing quotes.
		{
			Input: []byte(`hello`),
		},
		{
			Input: []byte(`"hello`),
		},
		{
			Input: []byte(`hello"`),
		},
		{
			Input: []byte(``),
		},
		{
			Input: []byte(`"`),
		},



		// Invalid escape sequences.
		{
			Input: []byte(`"bad\x"`),
		},
		{
			Input: []byte(`"bad\a"`),
		},
		{
			// Trailing backslash.
			Input: []byte(`"bad\"`),
		},



		// Invalid \uXXXX escapes.
		{
			// Too short.
			Input: []byte(`"\u00"`),
		},
		{
			// Invalid hex character.
			Input: []byte(`"\u00GG"`),
		},



		// Invalid surrogate pairs.
		{
			// High surrogate without low surrogate.
			Input: []byte(`"\uD800"`),
		},
		{
			// High surrogate followed by non-surrogate.
			Input: []byte(`"\uD800\u0041"`),
		},
		{
			// Lone low surrogate.
			Input: []byte(`"\uDC00"`),
		},
	}

	for testNumber, test := range tests {
		_, err := unquoteString(test.Input)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("INPUT: %s", test.Input)
			continue
		}
	}
}
