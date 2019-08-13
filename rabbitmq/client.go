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
	return
}

func (mq *ClientRabbitMQ) ExchangeDeclare(ex string) (err error) {
	err = mq.channel.ExchangeDeclare(
		ex,      // name
		"topic", // type
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
	return
}

func (mq *ClientRabbitMQ) Print(messages chan []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			log.Printf(" [x] ERROR: %s", err.Error())
			return
		}
	}()
	for message := range messages {
		log.Printf(" [x] %s", message)
		panic("123456") // 测试panic
	}
}

func (mq *ClientRabbitMQ) Keepalive() {
	// fixme
	for {
		for i := 0; i < 3; i++ {
			time.Sleep(5 * time.Second)
			if err := Init(); err != nil {
				log.Printf("[ERROR] MQ: Connection recover failed for %d times, %v\n", i+1, err)
				continue
			}
			log.Printf("[INFO] MQ: Connection recover OK. Total try %d times\n", i+1)
			break
		}
		select {
		case <-mq.done:
			return
		case <-mq.notifyClose:
			log.Printf("关啦")
		}
	}
}
