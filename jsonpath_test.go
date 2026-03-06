package json

import (
	"testing"
)

func TestJsonPath(t *testing.T) {

	tests := []struct{
		Build    func(p *jsonPath)
		Expected string
	}{
		// Empty path.
		{
			Build:    func(p *jsonPath) {},
			Expected: "$",
		},



		// Single key.
		{
			Build: func(p *jsonPath) {
				p.pushKey("foo")
			},
			Expected: "$.foo",
		},



		// Single index.
		{
			Build: func(p *jsonPath) {
				p.pushIndex(0)
			},
			Expected: "$[0]",
		},



		// Nested keys.
		{
			Build: func(p *jsonPath) {
				p.pushKey("foo")
				p.pushKey("bar")
				p.pushKey("baz")
			},
			Expected: "$.foo.bar.baz",
		},



		// Key then index.
		{
			Build: func(p *jsonPath) {
				p.pushKey("items")
				p.pushIndex(2)
			},
			Expected: "$.items[2]",
		},



		// Key, index, key.
		{
			Build: func(p *jsonPath) {
				p.pushKey("foo")
				p.pushIndex(3)
				p.pushKey("bar")
			},
			Expected: "$.foo[3].bar",
		},



		// Deep nested path.
		{
			Build: func(p *jsonPath) {
				p.pushKey("a")
				p.pushKey("b")
				p.pushIndex(0)
				p.pushKey("c")
				p.pushIndex(1)
			},
			Expected: "$.a.b[0].c[1]",
		},



		// Large index.
		{
			Build: func(p *jsonPath) {
				p.pushKey("items")
				p.pushIndex(999)
			},
			Expected: "$.items[999]",
		},
	}

	for testNumber, test := range tests {
		var path jsonPath
		test.Build(&path)
		actual := path.String()

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestJsonPath_pop(t *testing.T) {

	tests := []struct{
		Build    func(p *jsonPath)
		Expected string
	}{
		// Pop back to root.
		{
			Build: func(p *jsonPath) {
				p.pushKey("foo")
				p.pop()
			},
			Expected: "$",
		},



		// Pop last segment only.
		{
			Build: func(p *jsonPath) {
				p.pushKey("foo")
				p.pushKey("bar")
				p.pop()
			},
			Expected: "$.foo",
		},



		// Pop index, keep key.
		{
			Build: func(p *jsonPath) {
				p.pushKey("items")
				p.pushIndex(5)
				p.pop()
			},
			Expected: "$.items",
		},



		// Push, pop, push different.
		{
			Build: func(p *jsonPath) {
				p.pushKey("first")
				p.pop()
				p.pushKey("second")
			},
			Expected: "$.second",
		},



		// Simulate iterating array elements.
		{
			Build: func(p *jsonPath) {
				p.pushKey("items")
				p.pushIndex(0)
				p.pop()
				p.pushIndex(1)
				p.pop()
				p.pushIndex(2)
			},
			Expected: "$.items[2]",
		},



		// Pop on empty path is safe.
		{
			Build: func(p *jsonPath) {
				p.pop()
			},
			Expected: "$",
		},
		{
			Build: func(p *jsonPath) {
				p.pop()
				p.pop()
				p.pop()
			},
			Expected: "$",
		},
	}

	for testNumber, test := range tests {
		var path jsonPath
		test.Build(&path)
		actual := path.String()

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
