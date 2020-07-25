package sqlite

import (
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/k1LoW/regexq/parser"
)

const defaultTableName = "lines"

type Sqlite struct {
	tableName string
}

var (
	tmplCreateTable = template.Must(template.New("schema").Parse(`
CREATE TABLE IF NOT EXISTS {{.TableName}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,{{range $i, $value := .Schema}}
  {{$value}} TEXT,{{end}}
  created NUMERIC NOT NULL
);
{{range $i, $value := .Schema}}CREATE INDEX {{$.TableName}}_{{$value}}_idx ON {{$.TableName}}({{$value}});
{{end}}`))
	tmplInsert = template.Must(template.New("insert").Funcs(Funcs()).Parse(`INSERT INTO {{.TableName}}({{range $i, $value := .Schema}}{{$value}}, {{end}}created) VALUES ({{range $i, $value := .Schema}}'{{index $.In $value | escape_str}}', {{end}}datetime('now'));
`))
)

func New() *Sqlite {
	tableName := defaultTableName
	if n := os.Getenv("REGEXQ_TABLE_NAME"); n != "" {
		tableName = n
	}
	return &Sqlite{
		tableName: tableName,
	}
}

func (s *Sqlite) WriteSchema(w io.Writer, schema parser.Schema) error {
	params := map[string]interface{}{
		"TableName": s.tableName,
		"Schema":    schema,
	}
	return tmplCreateTable.Execute(w, params)
}

func (s *Sqlite) Write(w io.Writer, schema parser.Schema, in parser.Parsed) error {
	params := map[string]interface{}{
		"TableName": s.tableName,
		"In":        in,
		"Schema":    schema,
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
