package handlers

import (
	"encoding/json"
	"log"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

type ScanHandler struct {
	conf        *config.Config
	qConn       *amqp091.Connection
	scanService *services.ScanService
}

func NewScanHandler(conf *config.Config, qConn *amqp091.Connection, scanService *services.ScanService) *ScanHandler {
	return &ScanHandler{
		conf:        conf,
		qConn:       qConn,
		scanService: scanService,
	}
}

func (h *ScanHandler) GetTask() {
	// Connect to message queue
	ch, err := h.qConn.Channel()
	failOnError(constants.ErrorRabbitMQOpenChannel, err)
	defer ch.Close() // nolint

	queueName := h.conf.RabbitMQ.Queue.Name
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(constants.ErrorRabbitMQDeclareQueue, err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(constants.ErrorRabbitMQConsume, err)

	// Start consumer
	var forever chan struct{}

	go func() {
		for d := range msgs {
			h.handleTask(d.Body)

			if err := d.Ack(false); err != nil {
				failOnError(constants.ErrorRabbitMQAcknowledge, err)
			}
		}
	}()

	log.Printf(constants.MessageRabbitMQWaitingPrompt)
	<-forever
}

func (h *ScanHandler) handleTask(body []byte) {
	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		h.handleErrorTask(body, err)

		return
	}

	repositoryID, err := uuid.Parse(task.RepositoryIDStr)
	if err != nil {
		h.handleErrorTask(body, err)

		return
	}
	task.RepositoryID = repositoryID

	scanID, err := uuid.Parse(task.ScanIDStr)
	if err != nil {
		h.handleErrorTask(body, err)

		return
	}
	task.ScanID = scanID

	if err := h.scanService.RunTask(&task); err != nil {
		h.handleErrorTask(body, err)
	}
}

func (h *ScanHandler) handleErrorTask(body []byte, err error) {
	if err = h.scanService.RunErrorTask(body, err); err != nil {
		log.Panicf("%s: %s", constants.ErrorTaskFailed, err)
	}
}
