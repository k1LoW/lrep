package parser

import (
	"fmt"
	"os"
	"regexp"
)

const defaultRawKey = "_raw"

type Schema []string

type Parsed map[string]string

type Parser struct {
	re     *regexp.Regexp
	schema Schema
	rawKey string
}

// New return Parser
func New(regex string) *Parser {
	rawKey := defaultRawKey
	if k := os.Getenv("REGEXQ_RAW_KEY"); k != "" {
		rawKey = k
	}
	re := regexp.MustCompile(regex)
	schema := re.SubexpNames()
	for i := range schema {
		if schema[i] == "" {
			schema[i] = fmt.Sprintf("m%d", i)
		}
	}
	schema = append(schema, rawKey)
	return &Parser{
		re:     re,
		schema: schema,
		rawKey: rawKey,
	}
}

// Parse string
func (p *Parser) Parse(in string) Parsed {
	m := p.re.FindStringSubmatch(in)
	psd := Parsed(make(map[string]string, len(p.schema)))
	if len(m) > 0 {
		for i, v := range m {
			psd[p.schema[i]] = v
		}
	}
	psd[p.rawKey] = in
	return psd
}

// Schema return schema
func (p *Parser) Schema() Schema {
	return p.schema
}
