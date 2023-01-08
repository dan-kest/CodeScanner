package handlers

import (
	"github.com/dan-kest/cscanner/internal/handlers/payloads"
	"github.com/gofiber/fiber/v2"
)

func SendError(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(payloads.GenericResponse{
		Status:  "ERROR",
		Message: &message,
	})
}
