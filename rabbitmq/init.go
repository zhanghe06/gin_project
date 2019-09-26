package rabbitmq

import (
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

	//go MQ.Keepalive()

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

	return
}

func Close() {
	_ = MQ.channel.Close()
	_ = MQ.conn.Close()
}
