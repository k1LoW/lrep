package tsv

import (
	"io"

	"github.com/k1LoW/lrep/parser"
)

type TSV struct {
	w io.Writer
}

func New(w io.Writer) *TSV {
	return &TSV{
		w: w,
	}
}

func (l *TSV) WriteSchema(schema parser.Schema) error {
	first := true
	for _, k := range schema.Keys {
		if !first {
			if _, err := l.w.Write([]byte("\t")); err != nil {
				return err
			}
		}
		if _, err := l.w.Write([]byte(k)); err != nil {
			return err
		}
		first = false
	}
	if _, err := l.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

func (l *TSV) Write(schema parser.Schema, in parser.Parsed) error {
	first := true
	for _, k := range schema.Keys {
		if !first {
			if _, err := l.w.Write([]byte("\t")); err != nil {
				return err
			}
		}
		if _, err := l.w.Write([]byte(in.KVs[k])); err != nil {
			return err
		}
		first = false
	}
	if _, err := l.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
