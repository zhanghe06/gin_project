# LogRus

https://github.com/sirupsen/logrus

```
go get -u github.com/sirupsen/logrus
```

标准库: `log`没有封装

## 示例


- 标准库日志示例
```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	_, err := os.Open("app.log")
	if err != nil {
		fmt.Println(err)
	}

	logFile,err:=os.OpenFile("app.log",os.O_RDWR|os.O_CREATE,0)
	if err!=nil{
		log.Fatalln("读取日志文件失败",err)
	}
	defer logFile.Close()
	logger:=log.New(logFile,"\r\n",log.Ldate|log.Ltime)
	logger.Println("测试")
}
```

- logrus示例 main.go
```go
package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I'll be logged with common and other field")
	contextLogger.Warn("Me too")
}
```

## RotateLogs 日志切割

https://github.com/lestrrat-go/file-rotatelogs

```
go get -u github.com/lestrrat-go/file-rotatelogs
```
