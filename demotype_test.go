package json_test

import (
	"github.com/reiver/go-json"
)

type DemoType struct {
	value string
}

func DemoTypeOne() DemoType {
	return DemoType{
		value: "ONE",
	}
}

func DemoTypeTwo() DemoType {
	return DemoType{
		value: "TWO",
	}
}

func (receiver DemoType) MarshalJSON() ([]byte, error) {
	if "" == receiver.value {
		return nil, json.ErrEmpty("it's empty")
	}

	return []byte(`{"message":"`+receiver.value+`","note":"bing bong bang"}`), nil
}
