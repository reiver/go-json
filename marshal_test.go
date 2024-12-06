package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
//	"github.com/reiver/go-nul"
//	"github.com/reiver/go-opt"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		// 0
		{
			Value:     struct{}{},
			Expected: `{}`,
		},



		// 1
		{
			Value:     false,
			Expected: "false",
		},
		// 2
		{
			Value:     true,
			Expected: "true",
		},



		// 3
		{
			Value: uint64(0),
			Expected:    "0",
		},
		// 4
		{
			Value: uint64(1),
			Expected:    "1",
		},
		// 5
		{
			Value: uint64(2),
			Expected:    "2",
		},
		// 6
		{
			Value: uint64(3),
			Expected:    "3",
		},
		// 7
		{
			Value: uint64(4),
			Expected:    "4",
		},
		// 8
		{
			Value: uint64(5),
			Expected:    "5",
		},

		// 9
		{
			Value: uint64(254),
			Expected:    "254",
		},
		// 10
		{
			Value: uint64(255),
			Expected:    "255",
		},
		// 11
		{
			Value: uint64(256),
			Expected:    "256",
		},




		// 12
		{
			Value:     "",
			Expected: `""`,
		},

		// 13
		{
			Value:     "once",
			Expected: `"once"`,
		},
		// 14
		{
			Value:     "twice",
			Expected: `"twice"`,
		},
		// 15
		{
			Value:     "thrice",
			Expected: `"thrice"`,
		},
		// 16
		{
			Value:     "fource",
			Expected: `"fource"`,
		},



		// 17
		{
			Value: struct {
				Apple  bool
				Banana int
				Cherry string
			}{
				Apple: true,
				Banana: -1,
				Cherry: "🙂",
			},
			Expected: `{"Apple":true,"Banana":-1,"Cherry":"🙂"}`,
		},
		// 18
		{
			Value: struct {
				Apple  bool   `json:"apple"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: true,
				Banana: -1,
				Cherry: "🙂",
			},
			Expected: `{"apple":true,"banana":-1,"cherry":"🙂"}`,
		},



		// 19
		{
			Value: struct {
				Apple  bool   `json:"-"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: true,
				Banana: -1,
				Cherry: "🙂",
			},
			Expected: `{"banana":-1,"cherry":"🙂"}`,
		},
		// 20
		{
			Value: struct {
				Apple  bool   `json:"-,"`
				Banana int    `json:"banana"`
				Cherry string `json:"cherry"`
			}{
				Apple: true,
				Banana: -1,
				Cherry: "🙂",
			},
			Expected: `{"-":true,"banana":-1,"cherry":"🙂"}`,
		},



		// 21
		{
			Value: struct {
				A any
				B any
				C any
			}{
				A: nil,
				B: "something",
				C: 5,
			},
			Expected: `{"A":null,"B":"something","C":5}`,
		},



		// 22
		{
			Value: struct {
				Once  bool    `json:"-,omitempty"`
				Twice int     `json:"twice,omitempty"`
				Thrice string `json:"thrice,omitempty"`
				Fource any    `json:"fource,omitempty"`
			}{
				Once: true,
				Twice: -1,
				Thrice: "🙂",
				Fource:"something",
			},
			Expected: `{"-":true,"twice":-1,"thrice":"🙂","fource":"something"}`,
		},
		// 23
		{
			Value: struct {
				Once  bool    `json:"-,omitempty"`
				Twice int     `json:"twice,omitempty"`
				Thrice string `json:"thrice,omitempty"`
				Fource any    `json:"fource,omitempty"`
			}{
				Once: false,
				Twice: -1,
				Thrice: "🙂",
				Fource:"something",
			},
			Expected: `{"twice":-1,"thrice":"🙂","fource":"something"}`,
		},
		// 24
		{
			Value: struct {
				Once  bool    `json:"-,omitempty"`
				Twice int     `json:"twice,omitempty"`
				Thrice string `json:"thrice,omitempty"`
				Fource any    `json:"fource,omitempty"`
			}{
				Once: true,
				Twice: 0,
				Thrice: "🙂",
				Fource:"something",
			},
			Expected: `{"-":true,"thrice":"🙂","fource":"something"}`,
		},
		// 25
		{
			Value: struct {
				Once  bool    `json:"-,omitempty"`
				Twice int     `json:"twice,omitempty"`
				Thrice string `json:"thrice,omitempty"`
				Fource any    `json:"fource,omitempty"`
			}{
				Once: true,
				Twice: -1,
				Thrice: "",
				Fource:"something",
			},
			Expected: `{"-":true,"twice":-1,"fource":"something"}`,
		},
		// 26
		{
			Value: struct {
				Once  bool    `json:"-,omitempty"`
				Twice int     `json:"twice,omitempty"`
				Thrice string `json:"thrice,omitempty"`
				Fource any    `json:"fource,omitempty"`
			}{
				Once: true,
				Twice: -1,
				Thrice: "🙂",
				Fource:nil,
			},
			Expected: `{"-":true,"twice":-1,"thrice":"🙂"}`,
		},



		// 27
		{
			Value: struct {
				One DemoType
				Two DemoType `json:",omitempty"`
			}{
				One: DemoTypeOne(),
				Two: DemoTypeTwo(),
			},
			Expected: `{`+
				`"One":{"message":"ONE","note":"bing bong bang"}`+
				`,`+
				`"Two":{"message":"TWO","note":"bing bong bang"}`+
			`}`,
		},
		// 28
		{
			Value: struct {
				One DemoType
				Two DemoType `json:",omitempty"`
			}{
				One: DemoTypeOne(),
			},
			Expected: `{`+
				`"One":{"message":"ONE","note":"bing bong bang"}`+
			`}`,
		},
		// 29
		{
			Value: struct {
				One DemoType
				Two DemoType `json:"two,omitempty"`
			}{
				One: DemoTypeOne(),
			},
			Expected: `{`+
				`"One":{"message":"ONE","note":"bing bong bang"}`+
			`}`,
		},



		// 30
		{
			Value: struct {
				First  int `json:"first,string"`
				Second int `json:",string"`
			}{
				First: 11,
				Second: 12,
			},
			Expected: `{`+
				`"first":"11"`+
				`,`+
				`"Second":"12"`+
			`}`,
		},



/*
		{
			Value: struct {
				First  opt.Optional[string] `json:"first"`
				Second opt.Optional[string] `json:"second,omitempty"`
				Third  opt.Optional[string] `json:",omitempty"`
			}{
				First:  opt.Something("one"),
				Second: opt.Something("TWO"),
				Third:  opt.Something("3"),
			},
			Expected: `{"first":"one","second":"TWO","Third":"3"}`,
		},
		{
			Value: struct {
				First  opt.Optional[string] `json:"first"`
				Second opt.Optional[string] `json:"second,omitempty"`
				Third  opt.Optional[string] `json:",omitempty"`
			}{
				First:  opt.Something("one"),
				Second: opt.Something("TWO"),
			},
			Expected: `{"first":"one","second":"TWO"}`,
		},
		{
			Value: struct {
				First  opt.Optional[string] `json:"first"`
				Second opt.Optional[string] `json:"second,omitempty"`
				Third  opt.Optional[string] `json:",omitempty"`
			}{
				First:  opt.Something("one"),
			},
			Expected: `{"first":"one"}`,
		},
*/



/*
		{
			Value: struct {
				First  nul.Nullable[string] `json:"first"`
				Second nul.Nullable[string] `json:"second,omitempty"`
				Third  nul.Nullable[string] `json:",omitempty"`
			}{
				First:  nul.Something("one"),
				Second: nul.Something("TWO"),
				Third:  nul.Something("3"),
			},
			Expected: `{"first":"one","second":"TWO","Third":"3"}`,
		},
		{
			Value: struct {
				First  nul.Nullable[string] `json:"first"`
				Second nul.Nullable[string] `json:"second,omitempty"`
				Third  nul.Nullable[string] `json:",omitempty"`
			}{
				First:  nul.Something("one"),
				Second: nul.Something("TWO"),
				Third:  nul.Null[string](),
			},
			Expected: `{"first":"one","second":"TWO","Third":null}`,
		},
		{
			Value: struct {
				First  nul.Nullable[string] `json:"first"`
				Second nul.Nullable[string] `json:"second,omitempty"`
				Third  nul.Nullable[string] `json:",omitempty"`
			}{
				First:  nul.Something("one"),
				Second: nul.Something("TWO"),
			},
			Expected: `{"first":"one","second":"TWO"}`,
		},
		{
			Value: struct {
				First  nul.Nullable[string] `json:"first"`
				Second nul.Nullable[string] `json:"second,omitempty"`
				Third  nul.Nullable[string] `json:",omitempty"`
			}{
				First:  nul.Something("one"),
				Second: nul.Null[string](),
			},
			Expected: `{"first":"one","second":null}`,
		},
		{
			Value: struct {
				First  nul.Nullable[string] `json:"first"`
				Second nul.Nullable[string] `json:"second,omitempty"`
				Third  nul.Nullable[string] `json:",omitempty"`
			}{
				First:  nul.Something("one"),
			},
			Expected: `{"first":"one"}`,
		},
*/








		{
			Value: []string{
				"once",
				"twice",
				"thrice",
				"fource",
			},
			Expected: `["once","twice","thrice","fource"]`,
		},



		{
			Value: struct {
				First  []string `json:"first"`
				Second []string `json:"second,omitempty"`
				Third  []string `json:"second,omitempty"`
			}{
				First: []string{
					"once",
					"twice",
					"thrice",
					"fource",
				},
				Second: []string{},
			},
			Expected: `{"first":["once","twice","thrice","fource"],"second":[]}`,
		},








/*
		{
			Value: []DemoType2(nil),
			Expected: `[]`,
		},
		{
			Value: []DemoType2{},
			Expected: `[]`,
		},
		{
			Value: []DemoType2{
				DemoType2{
					Apple:  opt.Something("ONE"),
					Banana: opt.Something("TWO"),
					Cherry: opt.Something("THREE"),
				},
			},
			Expected: `[{"apple":"ONE","banana":"TWO","cherry":"THREE"}]`,
		},
		{
			Value: []DemoType2{
				DemoType2{
					Apple:  opt.Something("ONE"),
					Banana: opt.Something("TWO"),
					Cherry: opt.Something("THREE"),
				},
				DemoType2{
					Apple:  opt.Something("1"),
					Banana: opt.Something("2"),
				},
			},
			Expected: `[{"apple":"ONE","banana":"TWO","cherry":"THREE"},{"apple":"1","banana":"2"}]`,
		},
		{
			Value: []DemoType2{
				DemoType2{
					Apple:  opt.Something("ONE"),
					Banana: opt.Something("TWO"),
					Cherry: opt.Something("THREE"),
				},
				DemoType2{
					Apple:  opt.Something("1"),
					Banana: opt.Something("2"),
				},
				DemoType2{
					Apple:  opt.Something("one"),
				},
			},
			Expected: `[{"apple":"ONE","banana":"TWO","cherry":"THREE"},{"apple":"1","banana":"2"},{"apple":"one"}]`,
		},
		{
			Value: []DemoType2{
				DemoType2{
					Apple:  opt.Something("ONE"),
					Banana: opt.Something("TWO"),
					Cherry: opt.Something("THREE"),
				},
				DemoType2{
					Apple:  opt.Something("1"),
					Banana: opt.Something("2"),
				},
				DemoType2{
					Apple:  opt.Something("one"),
				},
				DemoType2{},
			},
			Expected: `[{"apple":"ONE","banana":"TWO","cherry":"THREE"},{"apple":"1","banana":"2"},{"apple":"one"},{}]`,
		},
*/








		{
			Value: map[string]string{},

			Expected: `{}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[string]string{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},



		{
			Value: map[string]any{},

			Expected: `{}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[string]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},



		{
			Value: map[any]any{},

			Expected: `{}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
			},

			Expected: `{"ONCE":"1"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
			},

			Expected: `{"ONCE":"1","TWICE":"2"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
			},

			Expected: `{"ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
		{
			Value: map[any]any{
				"ONCE":"1",
				"TWICE":"2",
				"THRICE":"3",
				"FOURCE":"4",
			},

			Expected: `{"FOURCE":"4","ONCE":"1","THRICE":"3","TWICE":"2"}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual json-marshaled is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				t.Logf("VALUE-KIND: %s", reflect.TypeOf(test.Value).Kind())
				continue
			}
		}
	}
}
