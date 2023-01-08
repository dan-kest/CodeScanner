package main

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/http"
)

func main() {
	conf := config.Read()

	http.RunHTTPServer(conf)
}
