package queue

import (
	"fmt"
	"log"

	"github.com/dan-kest/cscanner/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func buildURL(config *config.RabbitMQ) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	)
}

func Connect(config *config.RabbitMQ) *amqp.Connection {
	url := buildURL(config)
	qConn, err := amqp.Dial(url)
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}

	return qConn
}
