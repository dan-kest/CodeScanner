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
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func RunHTTPServer(conf *config.Config, db *gorm.DB, qConn *amqp091.Connection) {
	app := fiber.New()
	app.Use(cors.New())

	initilize(conf, app, db, qConn)

	log.Fatal(app.Listen(":" + conf.App.Port))
}

func initilize(conf *config.Config, app *fiber.App, db *gorm.DB, qConn *amqp091.Connection) {
	// Health Check API
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})

	api := app.Group("/api")

	repoRepository := repositories.NewRepoRepository(conf, db)
	scanRepository := repositories.NewScanRepository(conf, db)
	repoService := services.NewRepoService(conf, repoRepository, scanRepository)
	repoHandler := handlers.NewRepoHandler(qConn, repoService)
	routers.RegisterRepoRoute(api, repoHandler)
}
