package main

import (
	"chat/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

const (
	port = "8080"
)

func startApp() {
	mux := routers()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Printf("Starting web server on port %s", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)

}
