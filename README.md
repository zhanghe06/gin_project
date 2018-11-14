# Gin 项目实例

[![Build Status](https://travis-ci.org/zhanghe06/gin_project.svg?branch=master)](https://travis-ci.org/zhanghe06/gin_project)
[![Coverage Status](https://coveralls.io/repos/github/zhanghe06/gin_project/badge.svg?branch=master)](https://coveralls.io/github/zhanghe06/gin_project?branch=master)
[![The MIT License](http://img.shields.io/badge/license-MIT-green.svg)](https://github.com/zhanghe06/gin_project/blob/master/LICENSE)

获取本项目
```
go get github.com/zhanghe06/gin_project
```


## Golang

https://golang.org


## Gin

https://github.com/gin-gonic/gin

```
go get -u github.com/gin-gonic/gin
```

## GoVendor

https://github.com/kardianos/govendor

```
go get -u github.com/kardianos/govendor
```

## LogRus

https://github.com/sirupsen/logrus

```
go get -u github.com/sirupsen/logrus
```

## GOrm

https://github.com/jinzhu/gorm

```
go get -u github.com/jinzhu/gorm
```

## XOrm & XOrm tools

依赖
```
go get -u github.com/go-sql-driver/mysql        # Mysql
go get -u github.com/ziutek/mymysql/godrv       # MyMysql
go get -u github.com/lib/pq                     # Postgres
go get -u github.com/mattn/go-sqlite3           # SQLite
go get -u github.com/denisenkom/go-mssqldb      # MSSQL(依赖civil)
```

```
go get -u github.com/go-xorm/xorm
go get -u github.com/go-xorm/cmd/xorm
```

`go-mssqldb`依赖`cloud.google.com/go/civil`
```
mkdir -p $GOPATH/src/cloud.google.com
git clone https://github.com/GoogleCloudPlatform/google-cloud-go.git $GOPATH/src/cloud.google.com/go
```

备份、恢复、创建模型操作参考`etc`目录下的shell脚本


## Go-MySQL-Driver

https://github.com/go-sql-driver/mysql

```
go get -u github.com/go-sql-driver/mysql
```

## UUID

https://github.com/satori/go.uuid

```
go get -u github.com/satori/go.uuid
```


## 附加依赖
```
go get -u github.com/mattn/go-sqlite3
```


## GOrm 扩展

https://github.com/go-gormigrate/gormigrate

https://github.com/qor/validations


## etcd

https://github.com/etcd-io/etcd/tree/master/clientv3

```
go get -u go.etcd.io/etcd/clientv3
```


## wrk - a HTTP benchmarking tool

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
