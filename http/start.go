package http

import (
	"log"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/handlers"
	"github.com/dan-kest/cscanner/internal/repositories"
	"github.com/dan-kest/cscanner/internal/routers"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunHTTPServer(conf *config.Config) {
	app := fiber.New()
	app.Use(cors.New())

	initilize(conf, app)

	log.Fatal(app.Listen(":" + conf.App.Port))
}

func initilize(conf *config.Config, app *fiber.App) {
	// Health Check API
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})

	api := app.Group("/api")

	repoRepository := repositories.NewRepoRepository(conf, nil)
	repoService := services.NewRepoService(conf, repoRepository)
	repoHandler := handlers.NewRepoHandler(repoService)
	routers.RegisterRepoRoute(api, repoHandler)
}
