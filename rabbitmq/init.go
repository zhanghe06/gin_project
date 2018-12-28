package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var MQClient *amqp.Connection

func Init() (err error) {
	if MQClient != nil {
		return
	}
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.ip"),
		viper.GetString("rabbitmq.port"),
	)
	MQClient, err = amqp.Dial(url)
	return
}

func Close() {
	MQClient.Close()
}
