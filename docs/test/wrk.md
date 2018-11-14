# WRK - a HTTP benchmarking tool

https://github.com/wg/wrk

```bash
brew install wrk
brew install openssl
xcode-select --install
```

测试
```
wrk -t12 -c400 -d30s http://0.0.0.0:8080
```

结果
```
Running 30s test @ http://0.0.0.0:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.02ms   22.06ms 221.80ms   60.49%
    Req/Sec     1.50k   334.63     2.76k    78.61%
  186568 requests in 10.42s, 43.95MB read
  Socket errors: connect 0, read 235, write 0, timeout 0
Requests/sec:  17896.82
Transfer/sec:      4.22MB
```

开启连接池，性能提升
