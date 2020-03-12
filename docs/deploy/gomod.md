# GoMod

参考: [https://segmentfault.com/a/1190000018398763](https://segmentfault.com/a/1190000018398763)

```
export GO111MODULE=on
go mod init newapp      # 可以手动增加依赖go.uber.org/atomic v1.4.0或者让go自动发现和维护，下面build中会自动发现依赖包
go mod tidy             # 根据go.mod文件来处理依赖关系
go build main.go        # vendor_test.go文件中增加了import "go.uber.org/zap"的语句，IDE提示报错，执行build后依赖包记录在go.mod中
go mod download         # 依赖包会自动下载到$GOPATH/pkg/mod，多个项目可以共享缓存的mod
go mod vendor           # 从mod中拷贝到项目的vendor目录下，这样IDE就可以识别了
```

常见问题: [https://blog.csdn.net/zzhongcy/article/details/97243826](https://blog.csdn.net/zzhongcy/article/details/97243826)


开启模块支持后（set GO111MODULE=on），并不能与$GOPATH共存
或者在非$GOPATH下创建项目

```
# 1.13 默认开启
export GO111MODULE=on
# 1.13 之后才支持多个地址，之前版本只支持一个
export GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct
# 1.13 开始支持，配置私有 module，不去校验 checksum
export GOPRIVATE=*.corp.example.com,rsc.io/private
```

unknown revision
升级git版本
git config --global url."git@github.com:".insteadOf "https://github.com/"
git config --global url."git@e.coding.net:".insteadOf "https://e.coding.net/"

一定要配置成`--global`

解决翻墙，go.mod 文件添加:
```
replace (
    golang.org/x/<name:sys> <tag:v0.3.0> => github.com/golang/<name:sys> <tag:v0.3.0>
    cloud.google.com/go => github.com/googleapis/google-cloud-go master
    
    google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190522204451-c2c4e71fbf69
    
    google.golang.org/grpc => github.com/grpc/grpc-go v1.21.0
    
    google.golang.org/appengine => github.com/golang/appengine v1.6.1-0.20190515044707-311d3c5cf937
    
)
```

