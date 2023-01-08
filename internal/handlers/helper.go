package handlers

import (
	"context"
	"log"
	"time"

	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/handlers/payloads"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Error response wrapper.
func sendError(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(payloads.GenericResponse{
		Status:  constants.ResponseStatusError,
		Message: &message,
	})
}

// Panic on error, for critical issue.
func failOnError(message string, err error) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

// Publish message wrapper.
func publishMessage(qConn *amqp.Connection, queueName string, body []byte, timeout int) error {
	ch, err := qConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: constants.MIMETypeApplicationJSON,
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
