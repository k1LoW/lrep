package ltsv

import (
	"io"

	"github.com/k1LoW/regexq/parser"
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
	for k, v := range in {
		if !first {
			if _, err := l.w.Write([]byte("\t")); err != nil {
				return err
			}
		}
		if _, err := l.w.Write([]byte(k + ":" + v)); err != nil {
			return err
		}
		first = false
	}
	if _, err := l.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
