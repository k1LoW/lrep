package parser

import (
	"fmt"
	"regexp"
)

const rawKey = "_raw"

type Schema []string

type Parsed map[string]string

type Parser struct {
	re     *regexp.Regexp
	schema Schema
}

// New return Parser
func New(regex string) *Parser {
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
	psd[rawKey] = in
	return psd
}

// Schema return schema
func (p *Parser) Schema() Schema {
	return p.schema
}
