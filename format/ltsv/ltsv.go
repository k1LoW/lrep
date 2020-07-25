package ltsv

import (
	"io"

	"github.com/k1LoW/lrep/parser"
)

type LTSV struct {
	w io.Writer
}

func New(w io.Writer) *LTSV {
	return &LTSV{
		w: w,
	}
}

func (l *LTSV) WriteSchema(schema parser.Schema) error {
	return nil
}

func (l *LTSV) Write(schema parser.Schema, in parser.Parsed) error {
	first := true
	for _, k := range schema {
		if !first {
			if _, err := l.w.Write([]byte("\t")); err != nil {
				return err
			}
		}
		if _, err := l.w.Write([]byte(k + ":" + in[k])); err != nil {
			return err
		}
		first = false
	}
	if _, err := l.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
