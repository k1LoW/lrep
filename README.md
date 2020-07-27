# lrep [![Build Status](https://github.com/k1LoW/lrep/workflows/build/badge.svg)](https://github.com/k1LoW/lrep/actions) [![GitHub release](https://img.shields.io/github/release/k1LoW/lrep.svg)](https://github.com/k1LoW/lrep/releases)

lrep = l/re/p = line regular expression parser

## Usage

`lrep` converts a single-line string into structured data by using the regular expression capture groups as fields.

``` console
$ tail -f /var/log/access.log | lrep '^(\S*) \S* \S* \[(.*)\] "(.*)" (\S*) (\S*)'
{"_raw":"100.21.169.226 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/electronics HTTP/1.1\" 200 114","m0":"100.21.169.226 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/electronics HTTP/1.1\" 200 114","m1":"100.21.169.226","m2":"25/Jul/2020:16:25:05 +0900","m3":"GET /category/electronics HTTP/1.1","m4":"200","m5":"114"}
{"_raw":"104.141.81.229 - - [25/Jul/2020:16:25:05 +0900] \"GET /item/office/1680 HTTP/1.1\" 200 49","m0":"104.141.81.229 - - [25/Jul/2020:16:25:05 +0900] \"GET /item/office/1680 HTTP/1.1\" 200 49","m1":"104.141.81.229","m2":"25/Jul/2020:16:25:05 +0900","m3":"GET /item/office/1680 HTTP/1.1","m4":"200","m5":"49"}
{"_raw":"132.189.225.189 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/office HTTP/1.1\" 200 97","m0":"132.189.225.189 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/office HTTP/1.1\" 200 97","m1":"132.189.225.189","m2":"25/Jul/2020:16:25:05 +0900","m3":"GET /category/office HTTP/1.1","m4":"200","m5":"97"}
{"_raw":"228.189.133.138 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/networking?from=10 HTTP/1.1\" 200 47","m0":"228.189.133.138 - - [25/Jul/2020:16:25:05 +0900] \"GET /category/networking?from=10 HTTP/1.1\" 200 47","m1":"228.189.133.138","m2":"25/Jul/2020:16:25:05 +0900","m3":"GET /category/networking?from=10 HTTP/1.1","m4":"200","m5":"47"}
{"_raw":"24.111.108.90 - - [25/Jul/2020:16:25:06 +0900] \"GET /category/office HTTP/1.1\" 200 134","m0":"24.111.108.90 - - [25/Jul/2020:16:25:06 +0900] \"GET /category/office HTTP/1.1\" 200 134","m1":"24.111.108.90","m2":"25/Jul/2020:16:25:06 +0900","m3":"GET /category/office HTTP/1.1","m4":"200","m5":"134"}
[...]
```

or

``` console
$ lrep -f /var/log/access.log '^(\S*) \S* \S* \[(.*)\] "(.*)" (\S*) (\S*)'
```

Structured data fields are named as follows

| field name | description |
| --- | --- |
| `m0` | regexp submatches[0] |
| `m1` | regexp submatches[1] |
| `m2` | regexp submatches[2] |
| ... | ... |
| `_raw` ( default ) | raw string |

### Use named capturing group

`lrep` also allows you to specify a field name by using a named capture group.

``` console
$ tail -f /var/log/access.log | lrep '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
{"_raw":"96.114.162.71 - - [25/Jul/2020:16:21:03 +0900] \"GET /category/software HTTP/1.1\" 200 118","bytes":"118","host":"96.114.162.71","m0":"96.114.162.71 - - [25/Jul/2020:16:21:03 +0900] \"GET /category/software HTTP/1.1\" 200 118","request":"GET /category/software HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:03 +0900"}
{"_raw":"200.51.158.140 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/electronics HTTP/1.1\" 200 72","bytes":"72","host":"200.51.158.140","m0":"200.51.158.140 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/electronics HTTP/1.1\" 200 72","request":"GET /category/electronics HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"212.225.52.180 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/software HTTP/1.1\" 200 107","bytes":"107","host":"212.225.52.180","m0":"212.225.52.180 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/software HTTP/1.1\" 200 107","request":"GET /category/software HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"144.180.97.62 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 42","bytes":"42","host":"144.180.97.62","m0":"144.180.97.62 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 42","request":"GET /category/games HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"64.114.180.212 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 92","bytes":"92","host":"64.114.180.212","m0":"64.114.180.212 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 92","request":"GET /category/games HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
[...]
```

| field name | description |
| --- | --- |
| `m0` | regexp submatches[0] |
| `host` | regexp submatches[1] |
| `time` | regexp submatches[2] |
| `request` | regexp submatches[3] |
| `status` | regexp submatches[4] |
| `bytes` | regexp submatches[5] |
| `_raw` ( default ) | raw string |

### Ignore `m0` and `_raw`

If you want to ignore `m0`, use `--no-m0`.
And if you want to ignore `_raw`, you can use `--no-raw`.

``` console
$ tail -f /var/log/access.log | lrep --no-m0 --no-raw '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
{"bytes":"118","host":"100.39.167.131","request":"GET /category/toys HTTP/1.1","status":"200","time":"25/Jul/2020:17:46:26 +0900"}
{"bytes":"70","host":"36.30.101.105","request":"GET /item/electronics/1293 HTTP/1.1","status":"200","time":"25/Jul/2020:17:46:26 +0900"}
{"bytes":"123","host":"212.87.25.78","request":"GET /category/software HTTP/1.1","status":"200","time":"25/Jul/2020:17:46:26 +0900"}
{"bytes":"76","host":"84.189.195.199","request":"GET /category/office HTTP/1.1","status":"200","time":"25/Jul/2020:17:46:27 +0900"}
{"bytes":"103","host":"164.78.219.152","request":"GET /item/electronics/1175 HTTP/1.1","status":"200","time":"25/Jul/2020:17:46:28 +0900"}
[...]
```

## Install

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export LREP_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/lrep/releases/download/v$LREP_VERSION/lrep_$LREP_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export LREP_VERSION=X.X.X
$ yum install https://github.com/k1LoW/lrep/releases/download/v$LREP_VERSION/lrep_$LREP_VERSION-1_amd64.rpm
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/lrep
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/lrep/releases)

**go get:**

```console
$ go get github.com/k1LoW/lrep
```

## Support output format

`lrep` supports some output formats.

**JSON (`json`):**

``` console
$ tail -f /var/log/access.log | lrep -t json --no-m0 --no-raw '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
{"bytes":"46","host":"200.84.44.206","request":"GET /item/office/1367 HTTP/1.1","status":"200","time":"25/Jul/2020:17:49:18 +0900"}
{"bytes":"83","host":"212.204.65.142","request":"GET /item/jewelry/2431 HTTP/1.1","status":"200","time":"25/Jul/2020:17:49:18 +0900"}
{"bytes":"127","host":"220.171.132.54","request":"GET /category/office HTTP/1.1","status":"200","time":"25/Jul/2020:17:49:18 +0900"}
{"bytes":"119","host":"128.186.50.227","request":"GET /category/electronics?from=10 HTTP/1.1","status":"200","time":"25/Jul/2020:17:49:18 +0900"}
{"bytes":"104","host":"84.93.135.20","request":"GET /category/electronics HTTP/1.1","status":"200","time":"25/Jul/2020:17:49:18 +0900"}
[...]
```

**LTSV (`ltsv`):**

``` console
$ tail -f /var/log/access.log | lrep -t ltsv --no-m0 --no-raw '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
host:184.126.44.127     time:25/Jul/2020:17:49:48 +0900 request:GET /category/music HTTP/1.1    status:200      bytes:98
host:92.201.62.149      time:25/Jul/2020:17:49:48 +0900 request:GET /category/computers HTTP/1.1        status:200      bytes:100
host:100.216.170.167    time:25/Jul/2020:17:49:48 +0900 request:GET /category/computers HTTP/1.1        status:200      bytes:101
host:124.111.164.28     time:25/Jul/2020:17:49:49 +0900 request:GET /category/networking HTTP/1.1       status:200      bytes:60
host:44.93.53.78        time:25/Jul/2020:17:49:49 +0900 request:GET /item/software/310 HTTP/1.1 status:200      bytes:107
[...]
```

**SQLite Query (`sqlite`):**

``` console
$ tail -f /var/log/access.log | lrep -t sqlite --no-m0 '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
CREATE TABLE IF NOT EXISTS lines (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  host TEXT,
  time TEXT,
  request TEXT,
  status TEXT,
  bytes TEXT,
  _raw TEXT,
  created NUMERIC NOT NULL
);
INSERT INTO lines(host, time, request, status, bytes, _raw, created) VALUES ('224.51.78.136', '25/Jul/2020:17:51:24 +0900', 'GET /category/books HTTP/1.1', '200', '130', '224.51.78.136 - - [25/Jul/2020:17:51:24 +0900] "GET /category/books HTTP/1.1" 200 130', datetime('now'));
INSERT INTO lines(host, time, request, status, bytes, _raw, created) VALUES ('152.114.184.75', '25/Jul/2020:17:51:25 +0900', 'GET /category/finance HTTP/1.1', '200', '56', '152.114.184.75 - - [25/Jul/2020:17:51:25 +0900] "GET /category/finance HTTP/1.1" 200 56', datetime('now'));
INSERT INTO lines(host, time, request, status, bytes, _raw, created) VALUES ('168.57.224.190', '25/Jul/2020:17:51:25 +0900', 'GET /category/games?from=10 HTTP/1.1', '200', '60', '168.57.224.190 - - [25/Jul/2020:17:51:25 +0900] "GET /category/games?from=10 HTTP/1.1" 200 60', datetime('now'));
INSERT INTO lines(host, time, request, status, bytes, _raw, created) VALUES ('108.132.195.150', '25/Jul/2020:17:51:25 +0900', 'GET /category/electronics HTTP/1.1', '200', '123', '108.132.195.150 - - [25/Jul/2020:17:51:25 +0900] "GET /category/electronics HTTP/1.1" 200 123', datetime('now'));
INSERT INTO lines(host, time, request, status, bytes, _raw, created) VALUES ('220.228.91.30', '25/Jul/2020:17:51:25 +0900', 'GET /item/sports/2868 HTTP/1.1', '200', '50', '220.228.91.30 - - [25/Jul/2020:17:51:25 +0900] "GET /item/sports/2868 HTTP/1.1" 200 50', datetime('now'));
[...]
```

The query can be passed directly to `sqlite3` command.

``` console
$ cat /var/log/access.log | lrep -t sqlite --common | sqlite3 lines.db
```

**CSV (`csv`):**

``` console
$ lrep -f /var/log/access.log -t csv --combined
host,ident,user,time,method,resource,proto,status,bytes,referer,agent
40.222.173.129,-,-,27/Jul/2020:23:40:41 +0900,GET,/category/software,HTTP/1.1,200,95,-,Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
212.39.190.208,-,-,27/Jul/2020:23:40:41 +0900,GET,/category/networking,HTTP/1.1,200,118,-,Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; YTB730; GTB7.2; EasyBits GO v1.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.
0.30729; Media Center PC 6.0; .NET4.0C)
128.216.47.99,-,-,27/Jul/2020:23:40:41 +0900,GET,/category/electronics,HTTP/1.1,200,50,/category/electronics,Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
212.183.71.55,-,-,27/Jul/2020:23:40:41 +0900,GET,/category/jewelry,HTTP/1.1,200,128,-,"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.46 Safari/535.11"
[...]
```

**TSV (`tsv`):**

``` console
$ lrep -f /var/log/access.log -t tsv --combined
host    ident   user    time    method  resource        proto   status  bytes   referer agent
52.138.203.201  -       -       27/Jul/2020:23:42:05 +0900      GET     /item/networking/778    HTTP/1.1        200     115     /category/electronics   Mozilla/5.0 (Windows NT 6.0; rv:10.0.1) Gecko/20100101 Firefox/10.0.1
128.156.177.44  -       -       27/Jul/2020:23:42:05 +0900      GET     /category/electronics   HTTP/1.1        200     110     -       Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.77 Safari/535.7
68.111.118.174  -       -       27/Jul/2020:23:42:05 +0900      GET     /category/books HTTP/1.1        200     60      /category/software?from=10      Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:9.0.1) Gecko/20100101 Firefox/9.0.1
108.24.27.92    -       -       27/Jul/2020:23:42:05 +0900      GET     /category/games?from=10 HTTP/1.1        200     87      /category/games Mozilla/5.0 (Windows NT 6.0; rv:10.0.1) Gecko/20100101 Firefox/10.0.1
192.150.55.131  -       -       27/Jul/2020:23:42:06 +0900      GET     /category/electronics   HTTP/1.1        200     135     -       Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)
[...]
```

## Built-in regexp patterns

`lrep` has some [built-in regexp patterns](https://github.com/k1LoW/lrep/blob/master/parser/builtin.go).

``` console
$ tail -f /var/log/access.log | lrep --combined
{"_raw":"96.207.52.179 - - [25/Jul/2020:17:32:09 +0900] \"GET /item/jewelry/1307 HTTP/1.1\" 200 112 \"/category/software\" \"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)\"","agent":"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)","bytes":"112","host":"96.207.52.179","ident":"-","m0":"96.207.52.179 - - [25/Jul/2020:17:32:09 +0900] \"GET /item/jewelry/1307 HTTP/1.1\" 200 112 \"/category/software\" \"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)\"","method":"GET","proto":"HTTP/1.1","referer":"/category/software","resource":"/item/jewelry/1307","status":"200","time":"25/Jul/2020:17:32:09 +0900","user":"-"}
{"_raw":"168.132.101.209 - - [25/Jul/2020:17:32:09 +0900] \"GET /category/books HTTP/1.1\" 200 127 \"-\" \"Mozilla/5.0 (Windows NT 6.0; rv:10.0.1) Gecko/20100101 Firefox/10.0.1\"","agent":"Mozilla/5.0 (Windows NT 6.0; rv:10.0.1) Gecko/20100101 Firefox/10.0.1","bytes":"127","host":"168.132.101.209","ident":"-","m0":"168.132.101.209 - - [25/Jul/2020:17:32:09 +0900] \"GET /category/books HTTP/1.1\" 200 127 \"-\" \"Mozilla/5.0 (Windows NT 6.0; rv:10.0.1) Gecko/20100101 Firefox/10.0.1\"","method":"GET","proto":"HTTP/1.1","referer":"-","resource":"/category/books","status":"200","time":"25/Jul/2020:17:32:09+0900","user":"-"}
{"_raw":"188.204.198.141 - - [25/Jul/2020:17:32:09 +0900] \"GET /item/software/4286 HTTP/1.1\" 200 85 \"/category/office\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:9.0.1) Gecko/20100101 Firefox/9.0.1\"","agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:9.0.1) Gecko/20100101 Firefox/9.0.1","bytes":"85","host":"188.204.198.141","ident":"-","m0":"188.204.198.141 - - [25/Jul/2020:17:32:09 +0900] \"GET /item/software/4286 HTTP/1.1\" 200 85 \"/category/office\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:9.0.1) Gecko/20100101 Firefox/9.0.1\"","method":"GET","proto":"HTTP/1.1","referer":"/category/office","resource":"/item/software/4286","status":"200","time":"25/Jul/2020:17:32:09 +0900","user":"-"}
[...]
```

You can check the built-in regxp patterns by `lrep builtin` command.

``` console
$ lrep builtin
common       Common Log Format
combined     Combined Log Format
postgresql   PostgreSQL log
$ lrep builtin common
NAME
       common -- Common Log Format

REGEXP
       ^(?P<host>\S*) (?P<ident>\S*) (?P<user>\S*) \[(?P<time>.*)\] "(?P<method>\S+)(?: +(?P<resource>\S*) +(?P<proto>\S*?))?" (?P<status>\S*) (?P<bytes>\S*)

SAMPLE
       152.120.218.99 - - [25/Jul/2020:12:25:54 +0900] "GET /category/books HTTP/1.1" 200 67
```
