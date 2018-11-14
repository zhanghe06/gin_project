# Viper

https://github.com/spf13/viper

## 示例

main.go
```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("local")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	
	viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
    	fmt.Println("Config file changed:", e.Name)
    })

	informationPublic := viper.GetBool("information.public")
	fmt.Println("information.public:", informationPublic)

	author := viper.GetString("author")
	fmt.Println("author:", author)

	interest := viper.GetStringSlice("information.interest")
	fmt.Println("information.interest:", interest)
}
```


config/local.yaml
```yaml
TimeStamp: "2018-07-16 10:23:19"
Author: "admin"
PassWd: "123456"
Information:
  Name: "Harry"
  Age: "26"
  Interest:
  - "Reading"
  - "Running"
  - "Ping Pong"
  Image: "/path/header.png"
  Public: false
```


## 说明

- Get(key string) : interface{}
- GetBool(key string) : bool
- GetFloat64(key string) : float64
- GetInt(key string) : int
- GetString(key string) : string
- GetStringMap(key string) : map[string]interface{}
- GetStringMapString(key string) : map[string]string
- GetStringSlice(key string) : []string
- GetTime(key string) : time.Time
- GetDuration(key string) : time.Duration
- IsSet(key string) : bool
- AllSettings() : map[string]interface{}
