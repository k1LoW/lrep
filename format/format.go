package format

import (
	"io"

	"github.com/k1LoW/regexq/parser"
)

type Formatter interface {
	WriteSchema(w io.Writer, schema parser.Schema) error
	Write(w io.Writer, schema parser.Schema, in parser.Parsed) error
}
