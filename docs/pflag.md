# PFlag

https://github.com/spf13/pflag

标准库: `flag`

Docker源码中使用了`pflag`

## 示例

### 标准库`flag`示例

main.go
```go
package main

import (
	"flag"
	"fmt"
)

func main() {

	dataPath := flag.String("D", "/home/data/", "Data path")
	logFile := flag.String("l", "/tmp/log.log", "log file")

	flag.Parse()

	cmd := flag.Arg(0)
	fmt.Printf("action   : %s\n", cmd)
	fmt.Printf("data path: %s\n", *dataPath)
	fmt.Printf("log file : %s\n", *logFile)

	fmt.Printf("there are %d non-flag input param\n", flag.NArg())
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}

}
```

### 第三方库`pflag`示例

main.go
```go

```
