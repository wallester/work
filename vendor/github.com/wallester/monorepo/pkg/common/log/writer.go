package log

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/juju/errors"
)

type Writer struct {
	Out          io.Writer
	Indent       bool
	ProcessEvent func(event map[string]any) (map[string]any, error)
}

func (w Writer) Write(inputBytes []byte) (int, error) {
	var event map[string]any
	decoder := json.NewDecoder(bytes.NewReader(inputBytes))
	decoder.UseNumber()
	if err := decoder.Decode(&event); err != nil {
		return 0, errors.Annotate(err, "decoding event failed")
	}

	var err error
	if w.ProcessEvent != nil {
		event, err = w.ProcessEvent(event)
		if err != nil {
			return 0, errors.Annotate(err, "processing event failed")
		}
	}

	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if w.Indent {
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(event); err != nil {
		return 0, errors.Annotate(err, "encoding event failed")
	}

	if _, err := buf.WriteTo(w.Out); err != nil {
		return 0, errors.Annotate(err, "writing output failed")
	}

	return len(inputBytes), nil
}
