package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhanghe06/gin_project/config"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/etcds"
	"github.com/zhanghe06/gin_project/logs"
	"github.com/zhanghe06/gin_project/routers"
)

func main() {
	// 初始化配置
	config.Init()
	config.Watch()

	// 初始化日志
	err := logs.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer logs.Close()

	// 初始化数据库
	err = dbs.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	// 初始化ETCD
	err = etcds.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer etcds.Close()

	// 初始化路由
	router := routers.Init()

	// 启动服务
	router.Run()
}
