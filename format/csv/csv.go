package csv

import (
	"fmt"
	"io"
	"strings"

	"github.com/k1LoW/lrep/parser"
)

type CSV struct {
	w io.Writer
}

func New(w io.Writer) *CSV {
	return &CSV{
		w: w,
	}
}

func (l *CSV) WriteSchema(schema parser.Schema) error {
	first := true
	for _, k := range schema.Keys {
		if !first {
			if _, err := l.w.Write([]byte(",")); err != nil {
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

func (l *CSV) Write(schema parser.Schema, in parser.Parsed) error {
	first := true
	for _, k := range schema.Keys {
		if !first {
			if _, err := l.w.Write([]byte(",")); err != nil {
				return err
			}
		}
		escaped := strings.ReplaceAll(in.KVs[k], `"`, `""`)
		if strings.ContainsAny(in.KVs[k], "\r\n,\"") {
			escaped = fmt.Sprintf(`"%s"`, escaped)
		}
		if _, err := l.w.Write([]byte(escaped)); err != nil {
			return err
		}
		first = false
	}
	if _, err := l.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
