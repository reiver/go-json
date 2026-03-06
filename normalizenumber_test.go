package json_test

import (
	"testing"

	"github.com/reiver/go-json"
)

func TestNormalizeNumberString(t *testing.T) {

	tests := []struct{
		Value    string
		Expected string
	}{
		{
			Value:    "",
			Expected: "0",
		},
		{
			Value:    "0",
			Expected: "0",
		},
		{
			Value:    "00",
			Expected: "0",
		},
		{
			Value:    "000",
			Expected: "0",
		},
		{
			Value:    "1",
			Expected: "1",
		},
		{
			Value:    "01",
			Expected: "1",
		},
		{
			Value:    "001",
			Expected: "1",
		},
		{
			Value:    "42",
			Expected: "42",
		},
		{
			Value:    "042",
			Expected: "42",
		},
		{
			Value:    "-1",
			Expected: "-1",
		},
		{
			Value:    "-01",
			Expected: "-1",
		},
		{
			Value:    "+1",
			Expected: "1",
		},
		{
			Value:    "+01",
			Expected: "1",
		},
		{
			Value:    "-0",
			Expected: "0",
		},
		{
			Value:    "+0",
			Expected: "0",
		},
		{
			Value:    "3.14",
			Expected: "3.14",
		},
		{
			Value:    "3.140",
			Expected: "3.14",
		},
		{
			Value:    "3.1400",
			Expected: "3.14",
		},
		{
			Value:    "03.14",
			Expected: "3.14",
		},
		{
			Value:    "0.0",
			Expected: "0",
		},
		{
			Value:    "0.00",
			Expected: "0",
		},
		{
			Value:    "1.0",
			Expected: "1",
		},
		{
			Value:    "10.00",
			Expected: "10",
		},
		{
			Value:    "-3.14",
			Expected: "-3.14",
		},
		{
			Value:    "-0.0",
			Expected: "0",
		},



		{
			Value:    "100",
			Expected: "100",
		},

		{
			Value:   "0100",
			Expected: "100",
		},
		{
			Value:  "00100",
			Expected: "100",
		},
		{
			Value: "000100",
			Expected: "100",
		},

		{
			Value:    "100.",
			Expected: "100",
		},
		{
			Value:    "100.0",
			Expected: "100",
		},
		{
			Value:    "100.00",
			Expected: "100",
		},
		{
			Value:    "100.000",
			Expected: "100",
		},



		{
			Value:    "0.000037",
			Expected: "0.000037",
		},
		{
			Value:    "3.7e-5",
			Expected: "0.000037",
		},
		{
			Value:    "3.7E-5",
			Expected: "0.000037",
		},



		{
			Value:    "2800",
			Expected: "2800",
		},
		{
			Value:    "2.8e3",
			Expected: "2800",
		},
		{
			Value:    "2.8E3",
			Expected: "2800",
		},
	}

	for testNumber, test := range tests {
		actual := json.NormalizeNumberString(test.Value)

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual normalized value is not what was expected.", testNumber)
			t.Logf("VALUE:    %q", test.Value)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
