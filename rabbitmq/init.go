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

var Consumer = new(ClientRabbitMQ)

func Init() (err error) {
	//defer func() {
	//	if rec := recover(); rec != nil {
	//		err = fmt.Errorf("%v", rec)
	//		return
	//	}
	//}()
	err = Consumer.Connect()
	if err != nil {
		return
	}
	err = Consumer.Channel()
	if err != nil {
		return
	}
	Consumer.notifyClose = make(chan *amqp.Error)
	Consumer.channel.NotifyClose(Consumer.notifyClose) // channel connection 断开均通知
	Consumer.done = make(chan bool)

	//go Consumer.Keepalive()

	err = Consumer.ExchangeDeclare(Exchange)
	if err != nil {
		return
	}
	err = Consumer.QueueDeclare(QueueName)
	if err != nil {
		return
	}
	err = Consumer.QueueBinds(Exchange, BindingKeys)
	if err != nil {
		return
	}

	return
}

func Close() {
	_ = Consumer.channel.Close()
	_ = Consumer.conn.Close()
}
