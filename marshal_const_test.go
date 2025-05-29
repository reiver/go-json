package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

func TestMarshal_const(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
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
				t.Errorf("For test #%d, the actual json-marshaled value for the %T is not what was expected.", testNumber, test.Value)
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
