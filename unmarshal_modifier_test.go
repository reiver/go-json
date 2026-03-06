package json_test

import (
	"testing"

	"github.com/reiver/go-json"
)

func TestUnmarshalStringModifier(t *testing.T) {

	tests := []struct{
		JSON string
		Expected int
	}{
		{
			JSON:     `{"value":"123"}`,
			Expected: 123,
		},
		{
			JSON:     `{"value":"0"}`,
			Expected: 0,
		},
		{
			JSON:     `{"value":"-42"}`,
			Expected: -42,
		},
	}

	for testNumber, test := range tests {

		var dst struct {
			Value int `json:"value,string"`
		}

		err := json.Unmarshal([]byte(test.JSON), &dst)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: %s", test.JSON)
			continue
		}

		if test.Expected != dst.Value {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", dst.Value)
			t.Logf("JSON: %s", test.JSON)
			continue
		}
	}
}

func TestUnmarshalBareModifier(t *testing.T) {

	tests := []struct{
		JSON string
		Expected string
	}{
		{
			JSON:     `{"value":123}`,
			Expected: "123",
		},
		{
			JSON:     `{"value":true}`,
			Expected: "true",
		},
		{
			JSON:     `{"value":null}`,
			Expected: "null",
		},
	}

	for testNumber, test := range tests {

		var dst struct {
			Value string `json:"value,bare"`
		}

		err := json.Unmarshal([]byte(test.JSON), &dst)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: %s", test.JSON)
			continue
		}

		if test.Expected != dst.Value {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", dst.Value)
			t.Logf("JSON: %s", test.JSON)
			continue
		}
	}
}

func TestUnmarshalModifierNotReversible(t *testing.T) {

	var usher json.Usher

	usher.ImplantModifier("oneway", func(data []byte) ([]byte, error) {
		return append([]byte("X"), data...), nil
	}, nil)

	var dst struct {
		Value string `json:"value,oneway"`
	}

	err := usher.Unmarshal([]byte(`{"value":"hello"}`), &dst)
	if nil == err {
		t.Errorf("Expected an error but did not get one.")
		t.Logf("DST: %#v", dst)
		return
	}
}

func TestUnmarshalModifierRoundtrip(t *testing.T) {

	tests := []struct{
		Value int
	}{
		{
			Value: 42,
		},
		{
			Value: 0,
		},
		{
			Value: -99,
		},
	}

	type myStruct struct {
		Value int `json:"value,string"`
	}

	for testNumber, test := range tests {

		src := myStruct{Value: test.Value}

		marshaled, err := json.Marshal(src)
		if nil != err {
			t.Errorf("For test #%d, did not expect a marshal error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			continue
		}

		var dst myStruct
		err = json.Unmarshal(marshaled, &dst)
		if nil != err {
			t.Errorf("For test #%d, did not expect an unmarshal error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON: %s", marshaled)
			continue
		}

		if src.Value != dst.Value {
			t.Errorf("For test #%d, roundtrip value mismatch.", testNumber)
			t.Logf("EXPECTED: %#v", src.Value)
			t.Logf("ACTUAL:   %#v", dst.Value)
			t.Logf("JSON: %s", marshaled)
			continue
		}
	}
}
