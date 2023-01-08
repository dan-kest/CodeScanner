package main

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/http"
	"github.com/dan-kest/cscanner/queue"
)

func main() {
	conf := config.Read()
	dbConn := database.Connect(conf.Postgres)
	qConn := queue.Connect(conf.RabbitMQ)

	http.RunHTTPServer(conf, dbConn, qConn)
}
