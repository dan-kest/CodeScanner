package handlers

import (
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/handlers/payloads"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Error wrapper
func sendError(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(payloads.GenericResponse{
		Status:  constants.ResponseStatusError,
		Message: &message,
	})
}

func (h *RepoHandler) publish(task *models.Task) error {
	return nil
}
