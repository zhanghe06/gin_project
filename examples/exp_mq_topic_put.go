package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnErrorPut(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectionClosePut(conn *amqp.Connection) {
	_ = conn.Close()
}

func channelClosePut(ch *amqp.Channel) {
	_ = ch.Close()
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorPut(err, "Failed to connect to RabbitMQ")
	defer connectionClosePut(conn)

	ch, err := conn.Channel()
	failOnErrorPut(err, "Failed to open a channel")
	defer channelClosePut(ch)

	ex := "project.test.topic"
	//rk := "shanghai.huangpu"
	rk := "guangzhou.huangpu"
	//rk := "guangzhou.liwan"

	// 测试gin项目
	ex = "ex.project.topic"
	rk = "rk.project"

	err = ch.ExchangeDeclare(
		ex,      // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	failOnErrorPut(err, "Failed to declare an exchange")

	body := "test msg"
	err = ch.Publish(
		ex,    // exchange
		rk,    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnErrorPut(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s, ex: %s, rk: %s", body, ex, rk)
}
