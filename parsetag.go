package json

import (
	"strings"
)

// parseTag parses the json struct-field tags.
//
// For example:
//
//	struct {
//
//		// ...
//
//		Field  `json:"banana,omitempty"`
//		SkipMe `json:"-"`
//
//		// ...
//
//	}
func parseTag(tag string) (name string, skip bool, omitempty bool, modifiers []string) {
	 parts := strings.Split(tag, ",")
	if len(parts) <= 0 {
		return
	}

	if 1 == len(parts) && "-" == parts[0] {
		skip = true
		return
	}

	name, parts = parts[0], parts[1:]

	for _, part := range parts {
		switch part {
		case "omitempty":
			omitempty = true
		case "":
			// nothing here
		default:
			modifiers = append(modifiers, part)
		}
	}

	return
}
