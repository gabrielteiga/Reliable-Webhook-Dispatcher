package main

import (
	"fmt"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api"
)

const SERVER_PORT string = ":8080"

// @title Webhook Dispatcher
// @version 1.0
// @Description receive request from webhooks and deliver them async and reliable
// @schemes http
func main() {
	router := api.NewRouter()

	server := api.NewServer(SERVER_PORT, router)

	if err := server.Run(); err != nil {
		fmt.Println("something failed")
	}
}
