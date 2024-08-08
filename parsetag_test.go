package json

import (
	"testing"

	"slices"
)

func TestParseTag(t *testing.T) {

	tests := []struct{
		Tag string
		ExpectedName string
		ExpectedSkip bool
		ExpectedOmitEmpty bool
		ExpectedModifiers []string
	}{
		{
			Tag: "",
			ExpectedName: "",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},



		{
			Tag: "apple",
			ExpectedName: "apple",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "banana",
			ExpectedName: "banana",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "cherry",
			ExpectedName: "cherry",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},



		{
			Tag: "apple,",
			ExpectedName: "apple",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "banana,",
			ExpectedName: "banana",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "cherry,",
			ExpectedName: "cherry",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},



		{
			Tag: "apple,,",
			ExpectedName: "apple",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "banana,,",
			ExpectedName: "banana",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "cherry,,",
			ExpectedName: "cherry",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},



		{
			Tag: "-",
			ExpectedName: "",
			ExpectedSkip: true,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},
		{
			Tag: "-,",
			ExpectedName: "-",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: nil,
		},



		{
			Tag: ",omitempty",
			ExpectedName: "",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: nil,
		},
		{
			Tag: "abc,omitempty",
			ExpectedName: "abc",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: nil,
		},



		{
			Tag: "name,apple",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: []string{"apple"},
		},
		{
			Tag: "name,apple,banana",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: []string{"apple","banana"},
		},
		{
			Tag: "name,apple,banana,cherry",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: false,
			ExpectedModifiers: []string{"apple","banana","cherry"},
		},




		{
			Tag: "name,omitempty,apple,banana,cherry",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: []string{"apple","banana","cherry"},
		},
		{
			Tag: "name,apple,omitempty,banana,cherry",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: []string{"apple","banana","cherry"},
		},
		{
			Tag: "name,apple,banana,omitempty,cherry",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: []string{"apple","banana","cherry"},
		},
		{
			Tag: "name,apple,banana,cherry,omitempty",
			ExpectedName: "name",
			ExpectedSkip: false,
			ExpectedOmitEmpty: true,
			ExpectedModifiers: []string{"apple","banana","cherry"},
		},
	}

	for testNumber, test := range tests {

		actualName, actualSkip, actualOmitEmpty, actualModifiers := parseTag(test.Tag)

		{
			actual := actualName
			expected := test.ExpectedName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

		{
			actual := actualSkip
			expected := test.ExpectedSkip

			if expected != actual {
				t.Errorf("For test #%d, the actual 'skip' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				continue
			}
		}

		{
			actual := actualOmitEmpty
			expected := test.ExpectedOmitEmpty

			if expected != actual {
				t.Errorf("For test #%d, the actual 'omitempty' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				continue
			}
		}

		{
			actual := actualModifiers
			expected := test.ExpectedModifiers

			if !slices.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual 'modifiers' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}
	}
}
