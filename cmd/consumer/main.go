package main

import (
	"log"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/queue"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conf := config.Read()
	db := database.Connect(conf.Postgres)

	queue.InitConsumer(conf, db)
}
