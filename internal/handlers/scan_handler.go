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
	amqp "github.com/rabbitmq/amqp091-go"
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
	// ===== Connect to message queue =====

	ch, err := h.qConn.Channel()
	failOnError(constants.RabbitMQErrorOpenChannel, err)
	defer ch.Close()

	queueName := h.conf.RabbitMQ.Queue.Name
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(constants.RabbitMQErrorDeclareQueue, err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(constants.RabbitMQErrorConsume, err)

	// ===== Start consumer =====

	var forever chan struct{}

	go func() {
		for d := range msgs {
			h.handleTask(&d)
		}
	}()

	log.Printf(constants.RabbitMQWaitingPrompt)
	<-forever
}

func (h *ScanHandler) handleTask(d *amqp.Delivery) {
	var task *models.Task
	if err := json.Unmarshal(d.Body, task); err != nil {
		h.handleErrorTask(d.Body, err)

		return
	}

	repositoryID, err := uuid.Parse(task.RepositoryIDStr)
	if err != nil {
		h.handleErrorTask(d.Body, err)

		return
	}
	task.RepositoryID = repositoryID

	scanID, err := uuid.Parse(task.ScanIDStr)
	if err != nil {
		h.handleErrorTask(d.Body, err)

		return
	}
	task.ScanID = scanID

	h.scanService.RunTask(task)
}

func (h *ScanHandler) handleErrorTask(body []byte, err error) {
	if err = h.scanService.RunErrorTask(body, err); err != nil {
		log.Panicf("Error task failed: %s", err)
	}
}
