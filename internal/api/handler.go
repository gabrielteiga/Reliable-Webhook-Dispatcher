package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/gabrielteiga/webhook-dispatcher/docs"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/handler/health"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func PopulateHandlers(mux *http.ServeMux) {
	log.Printf("[INFO] - Populate api routes start")

	healthRoute(mux, "")
	swaggerRoute(mux, "")
	v1Routes(mux, "/api/v1")

	log.Printf("[INFO] - Populate api routes finished")
}

func swaggerRoute(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating swagger route")

	mux.Handle(
		createRoute(http.MethodGet, prefix, "/swagger/"),
		httpSwagger.WrapHandler,
	)
}

func healthRoute(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating health route")

	mux.Handle(
		createRoute(http.MethodGet, prefix, "/health"),
		middleware.HandlerLogger(http.HandlerFunc(health.Status)),
	)
}

func v1Routes(mux *http.ServeMux, prefix string) {
	log.Printf("[INFO] - Populating v1 routes")

	mux.Handle(
		createRoute(http.MethodGet, prefix, "/example"),
		middleware.HandlerLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("another example"))
		})),
	)
}

func createRoute(method, prefix, path string) string {
	if len(prefix) == 0 {
		if strings.HasPrefix(path, "/") {
			return fmt.Sprintf("%s %s", method, path)
		}
		return fmt.Sprintf("%s /%s", method, path)
	}

	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s %s%s", method, prefix, path)
	}

	return fmt.Sprintf("%s %s/%s", method, prefix, path)
}
