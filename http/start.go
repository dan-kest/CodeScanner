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
	"gorm.io/gorm"
)

func RunHTTPServer(conf *config.Config, db *gorm.DB) {
	app := fiber.New()
	app.Use(cors.New())

	initilize(conf, app, db)

	log.Fatal(app.Listen(":" + conf.App.Port))
}

func initilize(conf *config.Config, app *fiber.App, db *gorm.DB) {
	// Health Check API
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})

	api := app.Group("/api")

	repoRepository := repositories.NewRepoRepository(conf, db)
	repoService := services.NewRepoService(conf, repoRepository)
	repoHandler := handlers.NewRepoHandler(repoService)
	routers.RegisterRepoRoute(api, repoHandler)
}
