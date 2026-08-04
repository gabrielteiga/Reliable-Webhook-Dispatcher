package main

import (
	"github.com/gabrielteiga/webhook-dispatcher/internal/api"
)

// @title Webhook Dispatcher
// @version 1.0
// @Description receive request from webhooks and deliver them async and reliable
// @schemes http
func main() {
	api.Run()
}
