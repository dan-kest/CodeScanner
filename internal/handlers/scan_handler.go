package handlers

import (
	"encoding/json"
	"log"

	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/google/uuid"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ScanHandler struct {
	scanService *services.ScanService
}

func NewScanHandler(scanService *services.ScanService) *ScanHandler {
	return &ScanHandler{
		scanService: scanService,
	}
}

func (h *ScanHandler) HandleTask(d *amqp.Delivery) {
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
