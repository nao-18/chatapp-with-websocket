package main

import (
	"log"
	"net/http"

	"github.com/nao-18/chat-app-with-websocket/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8081")

	_ = http.ListenAndServe(":9000", mux)
}
