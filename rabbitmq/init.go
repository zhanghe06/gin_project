package rabbitmq

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const Exchange string = "ex.project.topic"
const RoutingKey string = "rk.project"
const QueueName string = "queue.project"

//const ConsumerTag string = "ctag-1"

var BindingKeys = []string{
	"rk.*",
}

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}

var MQ = new(ClientRabbitMQ)

func Init() (err error) {
	//defer func() {
	//	if rec := recover(); rec != nil {
	//		err = fmt.Errorf("%v", rec)
	//		return
	//	}
	//}()
	err = MQ.Connect()
	if err != nil {
		return
	}
	err = MQ.Channel()
	if err != nil {
		return
	}

	MQ.notifyClose = make(chan *amqp.Error)
	MQ.notifyConfirm = make(chan amqp.Confirmation)
	MQ.channel.NotifyClose(MQ.notifyClose)     // channel connection 断开均通知
	MQ.channel.NotifyPublish(MQ.notifyConfirm) // 消息发送确认通知
	MQ.done = make(chan bool)

	err = MQ.ExchangeDeclare(Exchange)
	if err != nil {
		return
	}
	err = MQ.QueueDeclare(QueueName)
	if err != nil {
		return
	}
	err = MQ.QueueBinds(Exchange, BindingKeys)
	if err != nil {
		return
	}

	// 消息处理
	MQ.messages = make(chan []byte)

	err = MQ.Consume(MQ.messages)
	if err != nil {
		return
	}

	// 仅仅打印消息（含异常处理，守护运行）
	i := 0
	go func() {
		for {
			i++
			MQ.Print()
			log.Printf(" [x] RabbitMQ Print Msg Retry: %d", i)
		}
	}()

	return
}

func Close() {
	_ = MQ.channel.Close()
	_ = MQ.conn.Close()
}
