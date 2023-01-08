package queue

import (
	"log"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/handlers"
	"github.com/dan-kest/cscanner/internal/repositories"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func InitConsumer(conf *config.Config, dbConn *gorm.DB, qConn *amqp091.Connection) {
	// ===== Connect to message queue =====

	ch, err := qConn.Channel()
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

	scanRepository := repositories.NewScanRepository(conf, dbConn)
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
