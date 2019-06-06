# Gin

[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

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

测试示例
```go
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Middleware01() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			fmt.Println("Middleware01 Defer")
		}()
		fmt.Println("Middleware01 Before")
		c.Next()
		fmt.Println("Middleware01 After")
	}
}
```

执行结果
```
Middleware01 Before
Middleware02 Before
Middleware03 Before

...

Middleware03 After
Middleware03 Defer
Middleware02 After
Middleware02 Defer
Middleware01 After
Middleware01 Defer
```

### Abort without return

1、在 Middleware02 中 c.Next() 前 Abort, 执行结果
```
Middleware01 Before
Middleware02 Before
Middleware02 After
Middleware02 Defer
Middleware01 After
Middleware01 Defer
```

2、在 Middleware02 中 c.Next() 后 Abort, 不会产生影响，正常执行

3、在 Middleware02 中 defer 里 Abort, 不会产生影响，正常执行


### Abort with return

1、在 Middleware02 中 c.Next() 前 Abort, 执行结果
```
Middleware01 Before
Middleware02 Before
Middleware02 Defer
Middleware01 After
Middleware01 Defer
```

2、在 Middleware02 中 c.Next() 后 Abort, 不会产生影响，正常执行

3、在 Middleware02 中 defer 里 Abort, 不会产生影响，正常执行


中间件使用`goroutines`注意:

https://github.com/gin-gonic/gin#goroutines-inside-a-middleware


## struct binding tag for Bind/ShouldBind

[go-playground/validator.v8](https://github.com/go-playground/validator)
[hdr-Baked_In_Validators_and_Tags](http://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags)
[gin_validator_v8_to_v9](https://github.com/go-playground/validator/tree/v9/_examples/gin-upgrading-overriding)

- required
- omitempty

### Bind Uri

https://github.com/gin-gonic/gin#bind-uri

```go
type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

var person Person
if err := c.ShouldBindUri(&person); err != nil {
    c.JSON(400, gin.H{"msg": err})
    return
}
```


## The difference between Bind and ShouldBind

Bind
```
It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
```

ShouldBind
```
Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid.
```

ShouldBind仅仅中止, 而Bind直接报400

现在 url path 还不支持绑定

## tag

- dive 用于切片
