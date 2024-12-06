package json

import (
	"github.com/reiver/go-erorr"
)

func (receiver *Usher) MergeAndMarshal(values ...any) ([]byte, error) {
	var length int = len(values)

	if length <= 0 {
		var empty = [2]byte{'{','}'}
		return empty[:], nil
	}

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var comma bool = false

	bytes = append(bytes, '{')

	for index, value := range values {
		result, err := receiver.Marshal(value)
		if nil != err {
			return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T): %w", 1+index, length, value, err)
		}

		// remove the beginning '{' and the ending '}'.
		{
			var lengthResult int = len(result)

			if lengthResult < 2 {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it is too short (length=%d) to be JSON objects — it is: %q", 1+index, length, value, lengthResult, result)
			}

			if '{' != result[0] {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it does not start with '{' — it is: %q", 1+index, length, value, result)
			}
			if '}' != result[lengthResult-1] {
				return nil, erorr.Errorf("json: problem merge-marshaling value №%d of %d (type: %T) — it does not end with '}' — it is: %q", 1+index, length, value, result)
			}

			result = result[1:lengthResult-1]

			// was this {}?, if "yes", then skip it
			if len(result) <= 0 {
				continue
			}
		}

		if comma {
			bytes = append(bytes, ',')
		}
		bytes = append(bytes, result...)
		comma = true
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
