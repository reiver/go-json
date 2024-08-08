package json

import (
	"bytes"
	gojson "encoding/json"
	"io"
)

func Compact(dst io.Writer, src []byte) error {
	var buffer bytes.Buffer

	err := gojson.Compact(&buffer, src)
	if nil != err {
		return err
	}

	_, err = buffer.WriteTo(dst)
	if nil != err {
		return err
	}

	return nil
}
