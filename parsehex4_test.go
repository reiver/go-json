package json

import (
	"testing"
)

func TestParseHex4(t *testing.T) {

	tests := []struct{
		Input    []byte
		Expected uint16
	}{
		// Zeros.
		{
			Input:    []byte("0000"),
			Expected: 0x0000,
		},



		// Single digits.
		{
			Input:    []byte("0001"),
			Expected: 0x0001,
		},
		{
			Input:    []byte("000a"),
			Expected: 0x000a,
		},
		{
			Input:    []byte("000A"),
			Expected: 0x000A,
		},
		{
			Input:    []byte("000f"),
			Expected: 0x000f,
		},
		{
			Input:    []byte("000F"),
			Expected: 0x000F,
		},



		// Common values.
		{
			Input:    []byte("0041"),
			Expected: 0x0041, // 'A'
		},
		{
			Input:    []byte("00e9"),
			Expected: 0x00e9, // 'é'
		},
		{
			Input:    []byte("00E9"),
			Expected: 0x00E9, // 'é' uppercase hex
		},
		{
			Input:    []byte("4e16"),
			Expected: 0x4e16, // '世'
		},
		{
			Input:    []byte("0020"),
			Expected: 0x0020, // space
		},



		// Surrogate values.
		{
			Input:    []byte("D83D"),
			Expected: 0xD83D,
		},
		{
			Input:    []byte("DE00"),
			Expected: 0xDE00,
		},
		{
			Input:    []byte("D800"),
			Expected: 0xD800,
		},
		{
			Input:    []byte("DFFF"),
			Expected: 0xDFFF,
		},



		// Maximum value.
		{
			Input:    []byte("FFFF"),
			Expected: 0xFFFF,
		},
		{
			Input:    []byte("ffff"),
			Expected: 0xFFFF,
		},



		// Mixed case.
		{
			Input:    []byte("aBcD"),
			Expected: 0xABCD,
		},
		{
			Input:    []byte("AbCd"),
			Expected: 0xABCD,
		},
	}

	for testNumber, test := range tests {
		actual, err := parseHex4(test.Input)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("INPUT: %s", test.Input)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: 0x%04X", test.Expected)
			t.Logf("ACTUAL:   0x%04X", actual)
			t.Logf("INPUT:    %s", test.Input)
			continue
		}
	}
}

func TestParseHex4_fail(t *testing.T) {

	tests := []struct{
		Input []byte
	}{
		// Invalid hex characters.
		{
			Input: []byte("000g"),
		},
		{
			Input: []byte("000G"),
		},
		{
			Input: []byte("00zz"),
		},
		{
			Input: []byte("GHIJ"),
		},
		{
			Input: []byte("00 0"),
		},
		{
			Input: []byte("00-1"),
		},
		{
			Input: []byte("00.0"),
		},
	}

	for testNumber, test := range tests {
		_, err := parseHex4(test.Input)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not get one.", testNumber)
			t.Logf("INPUT: %s", test.Input)
			continue
		}
	}
}
