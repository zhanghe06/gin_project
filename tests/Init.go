package tests

import (
	"fmt"
	"github.com/zhanghe06/gin_project/config"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/logs"
)

// 初始设置
func Setup() {
	// 初始化配置
	config.Init()
	config.Watch()

	// 初始化日志
	logs.Init()

	// 初始化数据库
	err := dbs.Init()
	if err != nil {
		fmt.Println(err)
	}
}


// 退出设置
func TearDown() {
	defer dbs.Close()
}
