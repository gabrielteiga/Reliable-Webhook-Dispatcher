package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/handler/health"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
)

const SERVER_PORT string = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.Handle(
		"GET /health",
		middleware.HandlerLogger(http.HandlerFunc(health.Get)),
	)

	if err := http.ListenAndServe(SERVER_PORT, mux); err != nil {
		fmt.Println("shutdown http server")
	}
}
