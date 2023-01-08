package queue

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/handlers"
	"github.com/dan-kest/cscanner/internal/repositories"
	"github.com/dan-kest/cscanner/internal/services"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func InitializeConsumer(conf *config.Config, dbConn *gorm.DB, qConn *amqp091.Connection) {
	scanRepository := repositories.NewScanRepository(conf, dbConn)
	scanService := services.NewScanService(conf, scanRepository)
	scanHandler := handlers.NewScanHandler(conf, qConn, scanService)

	scanHandler.GetTask()
}
