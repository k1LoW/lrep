package parser

type Builtin struct {
	Regexp   string
	TSKey    string
	TSFormat string
	Desc     string
	Samples  []string
}

var Builtins = map[string]Builtin{
	"common": Builtin{
		`^(?P<host>\S*) (?P<ident>\S*) (?P<user>\S*) \[(?P<time>.*)\] "(?P<method>\S+)(?: +(?P<resource>\S*) +(?P<proto>\S*?))?" (?P<status>\S*) (?P<bytes>\S*)`,
		"time",
		"%d/%b/%Y:%H:%M:%S %z",
		"Common Log Format",
		[]string{
			`152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67`,
		},
	},
	"combined": Builtin{
		`^(?P<host>\S*) (?P<ident>\S*) (?P<user>\S*) \[(?P<time>.*)\] "(?P<method>\S+)(?: +(?P<resource>\S*) +(?P<proto>\S*?))?" (?P<status>\S*) (?P<bytes>\S*) "(?P<referer>.*)" "(?P<agent>.*)"`,
		"time",
		"%d/%b/%Y:%H:%M:%S %z",
		"Combined Log Format",
		[]string{
			`152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67 "/item/electronics/4234" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)"`,
		},
	},
	"postgresql": Builtin{
		`^(?P<timestamp>.*) \[(?P<pid>\S*)\] (?P<message_type>[^:]*):\s*(?P<message>.*)$`,
		"timestamp",
		"2006-01-02 03:04:05.000 MST",
		"Postgresql log",
		[]string{
			`2020-07-25 05:37:40.021 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432`,
		},
	},
}
