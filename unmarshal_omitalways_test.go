package json_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-json"
)

type testOmitMe struct{}

func (testOmitMe) JSONOmitAlways() {}

var _ json.OmitAlways = testOmitMe{}

func TestUnmarshal_omitalways(t *testing.T) {

	type MyStruct struct {
		Name   string     `json:"name"`
		Secret testOmitMe `json:"secret"`
	}

	tests := []struct{
		Bytes    []byte
		Expected MyStruct
	}{
		{
			// OmitAlways field present in JSON — should be ignored.
			Bytes: []byte(`{"name":"hello","secret":{"anything":"here"}}`),
			Expected: MyStruct{
				Name: "hello",
			},
		},
		{
			// OmitAlways field absent from JSON — fine.
			Bytes: []byte(`{"name":"hello"}`),
			Expected: MyStruct{
				Name: "hello",
			},
		},
		{
			// OmitAlways field is null in JSON — should be ignored.
			Bytes: []byte(`{"name":"hello","secret":null}`),
			Expected: MyStruct{
				Name: "hello",
			},
		},
	}

	for testNumber, test := range tests {
		var actual MyStruct

		err := json.Unmarshal(test.Bytes, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", test.Expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}
	}
}
