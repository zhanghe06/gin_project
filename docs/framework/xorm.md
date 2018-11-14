# XOrm & XOrm tools

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
