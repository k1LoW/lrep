package sqlite

import (
	"io"
	"strings"
	"text/template"

	"github.com/k1LoW/regexq/parser"
)

type Sqlite struct{}

var (
	tmplCreateTable = template.Must(template.New("schema").Parse(`
CREATE TABLE IF NOT EXISTS logs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,{{range $i, $value := .Schema}}
  {{$value}} TEXT,{{end}}
  created NUMERIC NOT NULL
);
{{range $i, $value := .Schema}}CREATE INDEX logs_{{$value}}_idx ON logs({{$value}});
{{end}}`))
	tmplInsert = template.Must(template.New("insert").Funcs(Funcs()).Parse(`INSERT INTO logs({{range $i, $value := .Schema}}{{$value}}, {{end}}created) VALUES ({{range $i, $value := .Schema}}'{{index $.In $value | escape_str}}', {{end}}datetime('now'));
`))
)

func New() *Sqlite {
	return &Sqlite{}
}

func (s *Sqlite) WriteSchema(w io.Writer, schema parser.Schema) error {
	params := map[string]interface{}{
		"Schema": schema,
	}
	return tmplCreateTable.Execute(w, params)
}

func (s *Sqlite) Write(w io.Writer, schema parser.Schema, in parser.Parsed) error {
	params := map[string]interface{}{
		"In":     in,
		"Schema": schema,
	}
	return tmplInsert.Execute(w, params)
}

func Funcs() map[string]interface{} {
	return template.FuncMap{
		"escape_str": func(text string) string {
			r := strings.NewReplacer("'", "''")
			return r.Replace(text)
		},
	}
}
