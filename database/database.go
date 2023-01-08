package database

import (
	"context"
	"fmt"
	"time"

	"github.com/dan-kest/cscanner/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connectionTimeout int

func buildDSN(config *config.Postgres) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)
}

// Connect to database, return database instance.
func Connect(config *config.Postgres) *gorm.DB {
	logLevel := logger.Silent
	if config.IsPrintLog {
		logLevel = logger.Info
	}
	l := logger.Default.LogMode(logLevel)

	dsn := buildDSN(config)
	dialector := postgres.Open(dsn)
	options := &gorm.Config{
		Logger: l,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	connectionTimeout = config.ConnectionTimeout

	dbConn, err := gorm.Open(dialector, options)
	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	return dbConn
}

// Set instance to use connection timeout as configured.
func WithTimeout(db *gorm.DB) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(connectionTimeout))
	_ = cancel

	return db.WithContext(ctx)
}
