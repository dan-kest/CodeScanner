package routers

import (
	"github.com/dan-kest/cscanner/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRepoRoute(router fiber.Router, handler *handlers.RepoHandler) {
	cScanner := router.Group("/repo")

	cScanner.Get("/list", handler.ListRepo)
	cScanner.Get("/:id", handler.ViewRepo)
	cScanner.Post("/scan/:id", handler.ScanRepo)
	cScanner.Post("/", handler.CreateRepo)
	cScanner.Put("/:id", handler.UpdateRepo)
	cScanner.Delete("/:id", handler.DeleteRepo)
}
