package ltsv

import (
	"io"

	"github.com/k1LoW/regexq/parser"
)

type LTSV struct{}

func New() *LTSV {
	return &LTSV{}
}

func (j *LTSV) WriteSchema(w io.Writer, schema parser.Schema) error {
	return nil
}

func (j *LTSV) Write(w io.Writer, schema parser.Schema, in parser.Parsed) error {
	first := true
	for k, v := range in {
		if !first {
			if _, err := w.Write([]byte("\t")); err != nil {
				return err
			}
		}
		if _, err := w.Write([]byte(k + ":" + v)); err != nil {
			return err
		}
		first = false
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
