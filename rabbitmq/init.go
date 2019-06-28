package rabbitmq

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"time"
)

const Exchange string = "gin_project_exchange_topic"
const RoutingKey string = "gin_project.p"
const BindingKey string = "gin_project.*"
const QueueName string = "gin_project_queue_name"

//const ConsumerTag string = "ctag-1"

var MQConn *amqp.Connection
var MQChannel *amqp.Channel
var MQQueue amqp.Queue
var Done chan bool
var NotifyClose chan *amqp.Error

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}

func Init() (err error) {
	if MQConn != nil {
		return
	}
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.ip"),
		viper.GetString("rabbitmq.port"),
	)
	MQConn, err = amqp.Dial(url)
	return
}

func Close() {
	_ = MQChannel.Close()
	_ = MQConn.Close()
}

func Channel() (err error) {
	MQChannel, err = MQConn.Channel()
	if err != nil {
		return
	}
	NotifyClose = make(chan *amqp.Error)
	MQChannel.NotifyClose(NotifyClose)
	return
}

func ExchangeDeclare() (err error) {
	err = MQChannel.ExchangeDeclare(
		Exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	return
}

func QueueDeclare() (err error) {
	MQQueue, err = MQChannel.QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return

}

func QueueBind() (err error) {
	err = MQChannel.QueueBind(
		MQQueue.Name, // queue name
		BindingKey,   // routing key
		Exchange,     // exchange
		false,
		nil)
	return
}

func Consume() (mqMsg <-chan amqp.Delivery, err error) {
	mqMsg, err = MQChannel.Consume(
		MQQueue.Name, // queue
		"",           // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)

	go func() {
		for d := range mqMsg {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	return
}

func Publish(body string) (err error) {
	err = MQChannel.Publish(
		Exchange,   // exchange
		RoutingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			//ContentType: "application/json",
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return
}

func Keepalive() {
	// fixme
	for {
		for i := 0; i < 3; i++ {
			time.Sleep(5 * time.Second)
			if e := Init(); e != nil {
				log.Printf("[ERROR] MQ: Connection recover failed for %d times, %v\n", i+1, e)
				continue
			}
			log.Printf("[INFO] MQ: Connection recover OK. Total try %d times\n", i+1)
			break
		}
		select {
		case <-Done:
			return
		case <-NotifyClose:
			log.Printf("关啦")
		}
	}
}
