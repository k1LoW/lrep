package json

import (
	"encoding/json"
	"io"

	"github.com/k1LoW/regexq/parser"
)

type JSON struct{}

func New() *JSON {
	return &JSON{}
}

func (j *JSON) WriteSchema(w io.Writer, schema parser.Schema) error {
	return nil
}

func (j *JSON) Write(w io.Writer, schema parser.Schema, in parser.Parsed) error {
	e := json.NewEncoder(w)
	return e.Encode(in)
}
