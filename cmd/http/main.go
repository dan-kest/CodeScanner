package main

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/http"
)

func main() {
	conf := config.Read()
	db := database.Connect(conf.Postgres)

	http.RunHTTPServer(conf, db)
}
