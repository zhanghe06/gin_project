package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhanghe06/gin_project/config"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/logs"
	"github.com/zhanghe06/gin_project/rabbitmq"
	"github.com/zhanghe06/gin_project/routers"
	"github.com/zhanghe06/gin_project/validators"
)

func main() {
	// 初始化配置
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	config.Watch()

	// 初始化日志
	err = logs.Init()
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

	// 初始化消息队列
	err = rabbitmq.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()

	// 初始化ETCD
	//err = etcds.Init()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer etcds.Close()

	// 初始化路由
	router := routers.Init()

	// 注册校验器
	err = validators.Init()
	if err != nil {
		log.Fatal(err)
	}

	// 消息处理
	messages := make(chan []byte)

	err = rabbitmq.Consumer.Consume(messages)
	if err != nil {
		log.Fatal(err)
	}
	// 仅仅打印消息（含异常处理，守护运行）
	i := 0
	go func() {
		for {
			i++
			rabbitmq.Consumer.Print(messages)
			log.Printf(" [x] RabbitMQ Print Msg Retry: %d", i)
		}
	}()

	// 启动服务
	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
