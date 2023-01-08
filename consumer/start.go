package consumer

import (
	"fmt"
	"log"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/handlers"
	"github.com/dan-kest/cscanner/internal/repositories"
	"github.com/dan-kest/cscanner/internal/services"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
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

func InitConsumer(conf *config.Config, db *gorm.DB) {
	// ===== Connect to message queue =====

	url := buildURL(conf.RabbitMQ)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := conf.RabbitMQ.Queue.Name
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// ===== Init service =====

	scanRepository := repositories.NewScanRepository(conf, db)
	scanService := services.NewScanService(conf, scanRepository)
	scanHandler := handlers.NewScanHandler(scanService)

	// ===== Start consumer =====

	var forever chan struct{}

	go func() {
		for d := range msgs {
			scanHandler.HandleTask(&d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
