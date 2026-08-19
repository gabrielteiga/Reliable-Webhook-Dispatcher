package api

import (
	"log"
	"net/http"

	_ "github.com/gabrielteiga/webhook-dispatcher/docs"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/health"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/httputils"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/webhooks"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	log.Printf("[INFO] - Routes is beeing populate")

	swaggerRoute(mux, "")
	healthRoute(mux, "")
	v1Routes(mux, "/api/v1")

	log.Printf("[INFO] - Routes populated")

	return mux
}

func swaggerRoute(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating swagger route")

	mux.Handle(
		httputils.CreateRoute(http.MethodGet, prefix, "/swagger/"),
		httpSwagger.WrapHandler,
	)
}

func healthRoute(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating health route")

	health.RegisterRoutes(mux, prefix)
}

func v1Routes(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating v1 routes")

	webhooks.RegisterRoutes(mux, prefix)
}
