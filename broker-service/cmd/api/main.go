package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "80"
)

type Config struct{}

func main() {
	app := Config{}

	log.Println("Start broker service on post: ", webPort)

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := svr.ListenAndServe()

	if err != nil {
		log.Panic(err.Error())
	}
}
