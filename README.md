# Gin 项目实例

[![Build Status](https://travis-ci.org/zhanghe06/gin_project.svg?branch=master)](https://travis-ci.org/zhanghe06/gin_project)
[![Coverage Status](https://coveralls.io/repos/github/zhanghe06/gin_project/badge.svg?branch=master)](https://coveralls.io/github/zhanghe06/gin_project?branch=master)
[![The MIT License](http://img.shields.io/badge/license-MIT-green.svg)](https://github.com/zhanghe06/gin_project/blob/master/LICENSE)


项目演示
```bash
go get -u github.com/zhanghe06/gin_project
go get -u github.com/kardianos/govendor
govendor init
govendor sync
source env_local.sh
go run main.go
```

[http://localhost:8080/](http://localhost:8080/)


文档演示
```
cd docs
gitbook install
gitbook serve
```

[http://localhost:4000/](http://localhost:4000/)


## 语言哲学

换一种语言，意味着换一种信仰

语言不同，哲学不同，思维不同，最佳实践自然不同

[我为什么放弃Go语言](https://www.cnblogs.com/findumars/p/4097888.html)

[驳狗屎文 "我为什么放弃Go语言"](https://blog.csdn.net/cxlzxi/article/details/50284975)


## 常用语法特性

map 通过双赋值检测某个键存在
```
elem, ok = m[key]
```

指针操作符
```
& 变量取地址
* 指针取值
```

struct 定义json时，omitempty忽略零值和空值
```
Field int `json:"myName"`               // 以原始"myName"作为键名
Field int `json:"myName,omitempty"`     // 以原始"myName"作为键名，如果为空则忽略字段序列化
Field int `json:",omitempty"`           // 以原始"Field"作为键名，
Field int `json:"-"`                    // 忽略字段序列化
Field int `json:"-,"`                   // 以"-"作为键名
```

只有发送者才能关闭 channel，而不是接收者。向一个已经关闭的 channel 发送数据会引起 panic


## 调试

永远不要用`fmt.Print`, 应该使用`log.Info`（带时间和代码行数）
