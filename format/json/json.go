package json

import (
	"encoding/json"
	"io"

	"github.com/k1LoW/lrep/parser"
)

type JSON struct {
	w io.Writer
	e *json.Encoder
}

func New(w io.Writer) *JSON {
	return &JSON{
		w: w,
		e: json.NewEncoder(w),
	}
}

func (j *JSON) WriteSchema(schema parser.Schema) error {
	return nil
}

func (j *JSON) Write(schema parser.Schema, in parser.Parsed) error {
	return j.e.Encode(in)
}
