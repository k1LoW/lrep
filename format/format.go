package format

import (
	"github.com/k1LoW/lrep/parser"
)

type Formatter interface {
	WriteSchema(schema parser.Schema) error
	Write(schema parser.Schema, in parser.Parsed) error
}
