# lrep

l/re/p = line regular expression parser

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

### Support output format

`lrep` supports some output formats.

**JSON (`json`):**

``` console
$ tail -f /var/log/access.log | lrep '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
{"_raw":"96.114.162.71 - - [25/Jul/2020:16:21:03 +0900] \"GET /category/software HTTP/1.1\" 200 118","bytes":"118","host":"96.114.162.71","m0":"96.114.162.71 - - [25/Jul/2020:16:21:03 +0900] \"GET /category/software HTTP/1.1\" 200 118","request":"GET /category/software HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:03 +0900"}
{"_raw":"200.51.158.140 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/electronics HTTP/1.1\" 200 72","bytes":"72","host":"200.51.158.140","m0":"200.51.158.140 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/electronics HTTP/1.1\" 200 72","request":"GET /category/electronics HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"212.225.52.180 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/software HTTP/1.1\" 200 107","bytes":"107","host":"212.225.52.180","m0":"212.225.52.180 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/software HTTP/1.1\" 200 107","request":"GET /category/software HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"144.180.97.62 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 42","bytes":"42","host":"144.180.97.62","m0":"144.180.97.62 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 42","request":"GET /category/games HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
{"_raw":"64.114.180.212 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 92","bytes":"92","host":"64.114.180.212","m0":"64.114.180.212 - - [25/Jul/2020:16:21:04 +0900] \"GET /category/games HTTP/1.1\" 200 92","request":"GET /category/games HTTP/1.1","status":"200","time":"25/Jul/2020:16:21:04 +0900"}
[...]
```

**LTSV (`ltsv`):**

``` console
$ tail -f /var/log/access.log | lrep -t ltsv '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
m0:40.216.211.43 - - [25/Jul/2020:17:04:56 +0900] "GET /category/giftcards HTTP/1.1" 200 102    host:40.216.211.43      time:25/Jul/2020:17:04:56 +0900 request:GET /category/giftcards HTTP/1.1        status:200 bytes:102       _raw:40.216.211.43 - - [25/Jul/2020:17:04:56 +0900] "GET /category/giftcards HTTP/1.1" 200 102
m0:96.180.227.89 - - [25/Jul/2020:17:04:56 +0900] "GET /item/electronics/2212 HTTP/1.1" 200 48  host:96.180.227.89      time:25/Jul/2020:17:04:56 +0900 request:GET /item/electronics/2212 HTTP/1.1     status:200 bytes:48        _raw:96.180.227.89 - - [25/Jul/2020:17:04:56 +0900] "GET /item/electronics/2212 HTTP/1.1" 200 48
m0:144.123.171.183 - - [25/Jul/2020:17:04:56 +0900] "GET /item/software/2738 HTTP/1.1" 200 62   host:144.123.171.183    time:25/Jul/2020:17:04:56 +0900 request:GET /item/software/2738 HTTP/1.1        status:200 bytes:62        _raw:144.123.171.183 - - [25/Jul/2020:17:04:56 +0900] "GET /item/software/2738 HTTP/1.1" 200 62
m0:136.33.183.181 - - [25/Jul/2020:17:04:57 +0900] "GET /category/electronics?from=10 HTTP/1.1" 200 75  host:136.33.183.181     time:25/Jul/2020:17:04:57 +0900 request:GET /category/electronics?from=10 HTTP/1.1 status:200      bytes:75        _raw:136.33.183.181 - - [25/Jul/2020:17:04:57 +0900] "GET /category/electronics?from=10 HTTP/1.1" 200 75
m0:224.180.105.99 - - [25/Jul/2020:17:04:57 +0900] "GET /item/electronics/1226 HTTP/1.1" 200 106        host:224.180.105.99     time:25/Jul/2020:17:04:57 +0900 request:GET /item/electronics/1226 HTTP/1.status:200       bytes:106       _raw:224.180.105.99 - - [25/Jul/2020:17:04:57 +0900] "GET /item/electronics/1226 HTTP/1.1" 200 106
[...]
```

**SQLite Query (`sqlite`):**

``` console
$ tail -f /var/log/access.log | lrep -t sqlite '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)'
CREATE TABLE IF NOT EXISTS lines (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  m0 TEXT,
  host TEXT,
  time TEXT,
  request TEXT,
  status TEXT,
  bytes TEXT,
  _raw TEXT,
  created NUMERIC NOT NULL
);
CREATE INDEX lines_m0_idx ON lines(m0);
CREATE INDEX lines_host_idx ON lines(host);
CREATE INDEX lines_time_idx ON lines(time);
CREATE INDEX lines_request_idx ON lines(request);
CREATE INDEX lines_status_idx ON lines(status);
CREATE INDEX lines_bytes_idx ON lines(bytes);
CREATE INDEX lines__raw_idx ON lines(_raw);
INSERT INTO lines(m0, host, time, request, status, bytes, _raw, created) VALUES ('104.192.51.63 - - [25/Jul/2020:17:08:42 +0900] "GET /item/garden/2424 HTTP/1.1" 200 91', '104.192.51.63', '25/Jul/2020:17:08:42 +0900', 'GET /item/garden/2424 HTTP/1.1', '200', '91', '104.192.51.63 - - [25/Jul/2020:17:08:42 +0900] "GET /item/garden/2424 HTTP/1.1" 200 91', datetime('now'));
INSERT INTO lines(m0, host, time, request, status, bytes, _raw, created) VALUES ('64.135.129.49 - - [25/Jul/2020:17:08:43 +0900] "GET /item/giftcards/2049 HTTP/1.1" 200 43 "http://www.google.com/search?ie=UTF-8&q=google&sclient=psy-ab&q=Giftcards&oq=Giftcards&aq=f&aqi=g-vL1&aql=&pbx=1&bav=on.2,or.r_gc.r_pw.r_qf.,cf.osb&biw=2480&bih=349" "Mozilla/5.0 (compatible;', '64.135.129.49', '25/Jul/2020:17:08:43
+0900', 'GET /item/giftcards/2049 HTTP/1.1" 200 43 "http://www.google.com/search?ie=UTF-8&q=google&sclient=psy-ab&q=Giftcards&oq=Giftcards&aq=f&aqi=g-vL1&aql=&pbx=1&bav=on.2,or.r_gc.r_pw.r_qf.,cf.osb&biw=2480&bih=349', '"Mozilla/5.0', '(compatible;', '64.135.129.49 - - [25/Jul/2020:17:08:43 +0900] "GET /item/giftcards/2049 HTTP/1.1" 200 43 "http://www.google.com/search?ie=UTF-8&q=google&sclient=psy-ab&q=Giftcards&oq=Giftcards&aq=f&aqi=g-vL1&aql=&pbx=1&bav=on.2,or.r_gc.r_pw.r_qf.,cf.osb&biw=2480&bih=349" "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)"', datetime('now'));
INSERT INTO lines(m0, host, time, request, status, bytes, _raw, created) VALUES ('220.108.216.50 - - [25/Jul/2020:17:08:43 +0900] "GET /item/computers/1748 HTTP/1.1" 200 108', '220.108.216.50', '25/Jul/2020:17:08:43 +0900', 'GET /item/computers/1748 HTTP/1.1', '200', '108', '220.108.216.50 - - [25/Jul/2020:17:08:43 +0900] "GET /item/computers/1748 HTTP/1.1" 200 108', datetime('now'));
[...]
```

``` console
$ tail -f /var/log/access.log | lrep -t sqlite '^(?P<host>\S*) \S* \S* \[(?P<time>.*)\] "(?P<request>.*)" (?P<status>\S*) (?P<bytes>\S*)' | sqlite3 lines.db
```
