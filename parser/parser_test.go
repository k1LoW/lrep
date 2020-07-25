package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSchema(t *testing.T) {
	tests := []struct {
		regexp string
		rawKey string
		want   Schema
	}{
		{"(a+)b(c)", "", Schema([]string{"m0", "m1", "m2", "_raw"})},
		{"(?P<first>a+)b(c)", "", Schema([]string{"m0", "first", "m2", "_raw"})},
		{"(?P<first>a+)b(?P<last>c)", "", Schema([]string{"m0", "first", "last", "_raw"})},
		{"(a+)b(c)", "rawdata", Schema([]string{"m0", "m1", "m2", "rawdata"})},
	}
	for _, tt := range tests {
		os.Setenv("LREP_RAW_KEY", tt.rawKey)
		p := New(tt.regexp)
		got := p.Schema()
		if diff := cmp.Diff(got, tt.want, nil); diff != "" {
			t.Errorf("%s", diff)
		}
	}
	os.Setenv("LREP_RAW_KEY", "")
}

func TestParse(t *testing.T) {
	tests := []struct {
		regexp string
		line   string
		want   Parsed
	}{
		{
			// from https://docs.fluentd.org/parser/syslog
			regexp: `^\<(?P<pri>[0-9]+)\>(?P<time>[^ ]* {1,2}[^ ]* [^ ]*) (?P<host>[^ ]*) (?P<ident>[^ :\[]*)(?:\[(?P<pid>[0-9]+)\])?(?:[^\:]*\:)? *(?P<message>.*)$`,
			line:   "<6>Feb 28 12:00:00 192.168.0.1 fluentd[11111]: [error] Syslog test",
			want: Parsed(map[string]string{
				"m0":      "<6>Feb 28 12:00:00 192.168.0.1 fluentd[11111]: [error] Syslog test",
				"pri":     "6",
				"time":    "Feb 28 12:00:00",
				"host":    "192.168.0.1",
				"ident":   "fluentd",
				"pid":     "11111",
				"message": "[error] Syslog test",
				"_raw":    "<6>Feb 28 12:00:00 192.168.0.1 fluentd[11111]: [error] Syslog test",
			}),
		},
		{
			regexp: `^(?P<host>.*?) .*? .*? \[(?P<time>.*?)\] "(?P<method>\S+?)(?: +(?P<resource>.*?) +(?P<proto>\S*?))?" (?P<status>.*?) (?P<bytes>.*?) "(?P<referer>.*?)" "(?P<agent>.*?)"`,
			line:   `152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67 "/item/electronics/4234" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)"`,
			want: Parsed(map[string]string{
				"m0":       `152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67 "/item/electronics/4234" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)"`,
				"host":     "152.120.218.99",
				"time":     "25/Jul/2020:12:25:54 +0900",
				"method":   "GET",
				"resource": "/category/books",
				"proto":    "HTTP/1.1",
				"status":   "200",
				"bytes":    "67",
				"referer":  "/item/electronics/4234",
				"agent":    "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
				"_raw":     `152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67 "/item/electronics/4234" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)"`,
			}),
		},
	}
	for _, tt := range tests {
		p := New(tt.regexp)
		got := p.Parse(tt.line)
		if diff := cmp.Diff(got, tt.want, nil); diff != "" {
			t.Errorf("%s", diff)
		}
	}
}
