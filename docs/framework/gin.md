# Gin

https://github.com/gin-gonic/gin

```
go get -u github.com/gin-gonic/gin
```

## 示例

main.go
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```


### 单元测试

```bash
go test tests/*
```


## 中间件

在 Middleware02 中Abort, 执行结果
```
Middleware01 Before
Middleware02 Before
Middleware03 Before

Middleware01 After
```
