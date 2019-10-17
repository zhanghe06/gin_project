package rabbitmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"time"
)

type ClientRabbitMQ struct {
	uri           string
	conn          *amqp.Connection
	channel       *amqp.Channel
	queue         amqp.Queue
	messages 	  chan []byte
	done          chan bool
	notifyClose   chan *amqp.Error       // 异常关闭
	notifyConfirm chan amqp.Confirmation // 消息发送成功确认
}

func (mq *ClientRabbitMQ) Connect() (err error) {
	mq.uri = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.ip"),
		viper.GetString("rabbitmq.port"),
	)
	mq.conn, err = amqp.Dial(mq.uri)
	if err != nil {
		log.Printf("[amqp] connect error: %s\n", err)
		return
	}
	return
}

func (mq *ClientRabbitMQ) Channel() (err error) {
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		log.Printf("[amqp] get channel error: %s\n", err)
		return
	}
	// 通道设置为Confirm模式
	err = mq.channel.Confirm(false)
	if err != nil {
		log.Printf("[amqp] set channel confirm error: %s\n", err)
		return
	}
	return
}

func (mq *ClientRabbitMQ) ExchangeDeclare(ex string) (err error) {
	err = mq.channel.ExchangeDeclare(
		ex,      // name
		amqp.ExchangeTopic, // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	return
}

func (mq *ClientRabbitMQ) QueueDeclare(queueName string) (err error) {
	mq.queue, err = mq.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return

}

func (mq *ClientRabbitMQ) QueueBinds(ex string, bindingKeys []string) (err error) {
	for _, routingKey := range bindingKeys {
		err = mq.channel.QueueBind(
			mq.queue.Name, // queue name
			routingKey,    // routing key
			ex,            // exchange
			false,
			nil,
		)
		if err != nil {
			return
		}
	}
	return
}

//func (mq *ClientRabbitMQ) Consume() (mqMsg <-chan amqp.Delivery, err error) {
func (mq *ClientRabbitMQ) Consume(messages chan []byte) (err error) {
	var deliveries <-chan amqp.Delivery
	deliveries, err = mq.channel.Consume(
		mq.queue.Name, // queue
		"",            // consumer
		true,          // auto ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // args
	)
	if err != nil {
		return
	}

	// 将消息通过channel传递出去
	go func(deliveries <-chan amqp.Delivery, done chan bool, messages chan []byte) {
		for d := range deliveries {
			messages <- d.Body
		}
		done <- true
	}(deliveries, mq.done, messages)
	return
}

func (mq *ClientRabbitMQ) Publish(ex string, rk string, body string) (err error) {
	err = mq.channel.Publish(
		ex,    // exchange
		rk,    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			//ContentType: "application/json",
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return
	}

	// 消息确认超时时间
	confirmTimeout := time.After(time.Duration(10) * time.Second)

	select {
	case confirm := <-mq.notifyConfirm:
		if !confirm.Ack {
			//log.Printf(" [x] ERROR: publish confirm error")
			err = fmt.Errorf("%v", "publish confirm error")
		}
	case <-mq.notifyClose:
		err = fmt.Errorf("%v", "channel close")
	case <-confirmTimeout:
		err = fmt.Errorf("%v", "confirm timeout")
	}

	return
}

func (mq *ClientRabbitMQ) Print() {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			log.Printf(" [x] ERROR: %s", err.Error())
			return
		}
	}()
	for message := range mq.messages {
		log.Printf(" [x] %s", message)
		panic("123456") // 测试panic
	}
}

func (mq *ClientRabbitMQ) Keepalive() {
	// 断线重连
	c := 0
	for {
		select {
		case <-mq.done:
			log.Printf("[ERROR] MQ: Connection/Channel done")
			//return
		case <-mq.notifyClose:
			log.Printf("[ERROR] MQ: Connection/Channel closed")
		}
		c++
		time.Sleep(1 * time.Second)
		if err := Init(); err != nil {
			log.Printf("[ERROR] MQ: Connection failed for %d times, %v\n", c, err)
			continue
		}
		log.Printf("[INFO] MQ: Connection OK. Total try %d times\n", c)
		// 还原计数
		c = 0
	}
}
