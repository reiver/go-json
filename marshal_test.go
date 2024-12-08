package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
//	"github.com/reiver/go-nul"
//	"github.com/reiver/go-opt"
)

type OmitAlways1 struct {}
var _ json.OmitAlways = OmitAlways1{}
func (receiver OmitAlways1) JSONOmitAlways(){}

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
				Third  []string `json:"third,omitempty"`
			}{
				First: []string{
					"once",
					"twice",
					"thrice",
					"fource",
				},
				Second: []string{},
			},
			Expected: `{"first":["once","twice","thrice","fource"]}`,
		},



		{
			Value: struct {
				First  []string `json:"first"`
				Second []string `json:"second,omitempty"`
				Third  []string `json:"third,omitempty"`
			}{
				First: []string{
					"once",
					"twice",
					"thrice",
					"fource",
				},
				Second: []string{},
				Third:  []string{"hi"},
			},
			Expected: `{"first":["once","twice","thrice","fource"],"third":["hi"]}`,
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








		{
			Value: struct{
				Apple string       `json:"apple"`
				Banana OmitAlways1 `json:"banana"`
				Cherry int         `json:"cherry"`
			}{
				Apple: "one",
			},
			Expected: `{"apple":"one","cherry":0}`,
		},
		{
			Value: map[string]any{
				"apple":"one",
				"banana":OmitAlways1{},
				"cherry":5,
			},
			Expected: `{"apple":"one","cherry":5}`,
		},


		{
			Value: struct{
				Once     string `json:"once,omitempty"`
				Twice  []string `json:"twice,omitempty"`
				Thrice []string `json:"thrice,omitempty"`
				Fource []string `json:"fource,omitempty"`
				End      string `json:"end"`
			}{
				Once:   "not empty",
				Twice:  nil,
				Thrice: []string(nil),
				Fource: []string{},
				End: "here",
			},
			Expected: `{"once":"not empty","end":"here"}`,
		},









		{
			Value: struct{
				Something json.Const[bool] `json:"something" json.value:"false"`
			}{

			},
			Expected: `{"something":false}`,
		},
		{
			Value: struct{
				Something json.Const[bool] `json:"something" json.value:"true"`
			}{

			},
			Expected: `{"something":true}`,
		},



		{
			Value: struct{
				Something json.Const[int] `json:"something" json.value:"-5"`
			}{

			},
			Expected: `{"something":-5}`,
		},
		{
			Value: struct{
				Something json.Const[int8] `json:"something" json.value:"-5"`
			}{

			},
			Expected: `{"something":-5}`,
		},
		{
			Value: struct{
				Something json.Const[int16] `json:"something" json.value:"-5"`
			}{

			},
			Expected: `{"something":-5}`,
		},
		{
			Value: struct{
				Something json.Const[int32] `json:"something" json.value:"-5"`
			}{

			},
			Expected: `{"something":-5}`,
		},
		{
			Value: struct{
				Something json.Const[int64] `json:"something" json.value:"-5"`
			}{

			},
			Expected: `{"something":-5}`,
		},



		{
			Value: struct{
				Something json.Const[string] `json:"something" json.value:"apple banana cherry"`
				Here      string             `json:"here"`
				Else      string             `json:"else"`
			}{
				Here: "=-=-=",
				Else:                                                     "apple banana cherry",
			},
			Expected: `{"something":"apple banana cherry","here":"=-=-=","else":"apple banana cherry"}`,
		},
		{
			Value: struct{
				Something json.Const[string] `json:"something" json.value:"\tonce\ntwice\nthrice\nfource\n"`
				Here      string             `json:"here"`
				Else      string             `json:"else"`
			}{
				Here: "=-=-=",
				Else:                                                     "\tonce\ntwice\nthrice\nfource\n",
			},
			Expected: `{"something":"\tonce\ntwice\nthrice\nfource\n","here":"=-=-=","else":"\tonce\ntwice\nthrice\nfource\n"}`,
		},



		{
			Value: struct{
				Something json.Const[uint] `json:"something" json.value:"5"`
			}{

			},
			Expected: `{"something":5}`,
		},
		{
			Value: struct{
				Something json.Const[uint8] `json:"something" json.value:"5"`
			}{

			},
			Expected: `{"something":5}`,
		},
		{
			Value: struct{
				Something json.Const[uint16] `json:"something" json.value:"5"`
			}{

			},
			Expected: `{"something":5}`,
		},
		{
			Value: struct{
				Something json.Const[uint32] `json:"something" json.value:"5"`
			}{

			},
			Expected: `{"something":5}`,
		},
		{
			Value: struct{
				Something json.Const[uint64] `json:"something" json.value:"5"`
			}{

			},
			Expected: `{"something":5}`,
		},









		{
			Value: struct { // Manitoban
				GivenName               string  `json:"given-name"`
				AdditionalNames       []string  `json:"additional-names,omitempty"`
				FamilyName              string  `json:"family-name"`
				HomeCountry  json.Const[string] `json:"home-country"  json.value:"Canada"`
				HomeProvince json.Const[string] `json:"home-province" json.value:"Manitoba"`
				HomeCity                string  `json:"home-city"`
			}{
				GivenName:  "Joe",
				FamilyName: "Blow",
				HomeCity:   "The Pas",
			},
			Expected: `{"given-name":"Joe","family-name":"Blow","home-country":"Canada","home-province":"Manitoba","home-city":"The Pas"}`,
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
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				t.Logf("VALUE-KIND: %s", reflect.TypeOf(test.Value).Kind())
				t.Logf("VALUE-TYPE: %T", test.Value)
				continue
			}
		}
	}
}
