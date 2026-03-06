package json

import (
	"strconv"
	"strings"
)

// jsonPath tracks the current location in the JSON structure
// for error reporting. Produces paths like "$.foo.bar[2].baz".
type jsonPath struct {
	segments []string
}

func (p *jsonPath) pushKey(key string) {
	p.segments = append(p.segments, "."+key)
}

func (p *jsonPath) pushIndex(index int) {
	p.segments = append(p.segments, "["+strconv.Itoa(index)+"]")
}

func (p *jsonPath) pop() {
	if len(p.segments) > 0 {
		p.segments = p.segments[:len(p.segments)-1]
	}
}

func (p *jsonPath) String() string {
	var builder strings.Builder

	builder.WriteByte('$')

	for _, seg := range p.segments {
		builder.WriteString(seg)
	}

	return builder.String()
}
