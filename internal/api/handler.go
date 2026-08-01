package api

import (
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/handler/health"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
)

func PopulateHandlers(mux *http.ServeMux) {
	mux.Handle(
		"GET /health",
		middleware.HandlerLogger(http.HandlerFunc(health.Get)),
	)
}
