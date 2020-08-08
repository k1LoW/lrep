package parser

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/itchyny/timefmt-go"
)

const defaultRawKey = "_raw"

type Schema struct {
	Keys  []string
	TSKey string
}

type Parsed struct {
	KVs     map[string]string
	TSKey   string
	TSValue time.Time
}

type Parser struct {
	re       *regexp.Regexp
	schema   Schema
	rawKey   string
	tsParser func(layout, value string) (time.Time, error)
	tsKey    string
	tsFormat string

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

func TSKey(tsKey string) Option {
	return func(p *Parser) {
		p.tsKey = tsKey
	}
}

func TSFormat(tsFormat string) Option {
	return func(p *Parser) {
		p.tsFormat = tsFormat
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
	now := time.Now().Format("2006-01-02")
	if strings.Contains(p.tsFormat, "%") {
		p.tsFormat = fmt.Sprintf("%%F %s", p.tsFormat)
		p.tsParser = func(layout, value string) (time.Time, error) {
			return timefmt.Parse(fmt.Sprintf("%s %s", now, value), layout)
		}
	} else {
		p.tsFormat = fmt.Sprintf("2006-01-02 %s", p.tsFormat)
		p.tsParser = func(layout, value string) (time.Time, error) {
			return time.Parse(layout, fmt.Sprintf("%s %s", now, value))
		}
	}

	keys := re.SubexpNames()
	for i := range keys {
		if keys[i] == "" {
			keys[i] = fmt.Sprintf("m%d", i)
		}
	}
	if !p.noRaw {
		keys = append(keys, rawKey)
	}
	p.schema = Schema{
		Keys:  keys,
		TSKey: p.tsKey,
	}
	return p
}

// Parse string
func (p *Parser) Parse(in string) Parsed {
	m := p.re.FindStringSubmatch(in)
	psd := Parsed{
		KVs:   make(map[string]string, len(p.schema.Keys)),
		TSKey: p.tsKey,
	}
	if len(m) > 0 {
		for i, v := range m {
			if i == 0 && p.noM0 {
				continue
			}
			psd.KVs[p.schema.Keys[i]] = v
		}
	}
	if !p.noRaw {
		psd.KVs[p.rawKey] = in
	}
	if p.haveTS() {
		t, _ := p.tsParser(p.tsFormat, psd.KVs[p.tsKey])
		psd.TSValue = t
	}
	return psd
}

// Schema return schema
func (p *Parser) Schema() Schema {
	if p.noM0 {
		return Schema{
			Keys:  p.schema.Keys[1:],
			TSKey: p.schema.TSKey,
		}
	}
	return p.schema
}

func (p *Parser) haveTS() bool {
	return p.tsKey != ""
}
