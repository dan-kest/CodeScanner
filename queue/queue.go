package queue

import (
	"fmt"
	"log"

	"github.com/dan-kest/cscanner/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

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
	failOnError(err, "Failed to connect to RabbitMQ")

	return qConn
}
