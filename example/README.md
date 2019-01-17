
start test web
```bash
./bin/hacklang example/web.hack
```

ab result
```bash
$ ab -k -c 300 -n 10000 "http://127.0.0.1:8083/"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8083

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      300
Time taken for tests:   0.598 seconds
Complete requests:      10000
Failed requests:        0
Keep-Alive requests:    10000
Total transferred:      1360000 bytes
HTML transferred:       110000 bytes
Requests per second:    16713.27 [#/sec] (mean)
Time per request:       17.950 [ms] (mean)
Time per request:       0.060 [ms] (mean, across all concurrent requests)
Transfer rate:          2219.73 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   3.2      0      23
Processing:     0   17  11.8     14      93
Waiting:        0   17  11.8     14      93
Total:          0   18  12.0     15      93

Percentage of the requests served within a certain time (ms)
  50%     15
  66%     20
  75%     24
  80%     27
  90%     34
  95%     40
  98%     48
  99%     54
 100%     93 (longest request)
```