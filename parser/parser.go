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

	noM0  bool
	noRaw bool
}

// Option function change Parser option
type Option func(*Parser)

func NoM0() Option {
	return func(p *Parser) {
		p.noM0 = true
	}
}

func NoRaw() Option {
	return func(p *Parser) {
		p.noRaw = true
	}
}

// New return Parser
func New(regex string, opts ...Option) *Parser {
	rawKey := defaultRawKey
	if k := os.Getenv("LREP_RAW_KEY"); k != "" {
		rawKey = k
	}
	re := regexp.MustCompile(regex)
	p := &Parser{
		re:     re,
		rawKey: rawKey,
	}
	for _, opt := range opts {
		opt(p)
	}
	schema := re.SubexpNames()
	for i := range schema {
		if schema[i] == "" {
			schema[i] = fmt.Sprintf("m%d", i)
		}
	}
	if !p.noRaw {
		schema = append(schema, rawKey)
	}
	p.schema = schema
	return p
}

// Parse string
func (p *Parser) Parse(in string) Parsed {
	m := p.re.FindStringSubmatch(in)
	psd := Parsed(make(map[string]string, len(p.schema)))
	if len(m) > 0 {
		for i, v := range m {
			if i == 0 && p.noM0 {
				continue
			}
			psd[p.schema[i]] = v
		}
	}
	if !p.noRaw {
		psd[p.rawKey] = in
	}
	return psd
}

// Schema return schema
func (p *Parser) Schema() Schema {
	if p.noM0 {
		return p.schema[1:]
	}
	return p.schema
}
