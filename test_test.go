package json_test

type Inner struct {
	Name string `json:"name"`
}

type InnerOmit struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

type Inner2 struct {
	Color string `json:"color"`
}

type Outer struct {
	Inner
	Label string `json:"label"`
}
