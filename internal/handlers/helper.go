package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WrapErrorsss(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(ErrorResponse{
		Message: message,
	})
}
