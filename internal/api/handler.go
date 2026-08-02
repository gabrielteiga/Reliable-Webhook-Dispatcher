package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/handler/health"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
)

func PopulateHandlers(mux *http.ServeMux) {
	healthRoute(mux, "")
	v1Routes(mux, "/api/v1")
	// mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Routes(mux, "api/v1/")))
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

func healthRoute(mux *http.ServeMux, prefix string) {
	mux.Handle(
		createRoute(http.MethodGet, prefix, "/health"),
		middleware.HandlerLogger(http.HandlerFunc(health.Status)),
	)
}

func v1Routes(mux *http.ServeMux, prefix string) {
	mux.Handle(
		createRoute(http.MethodGet, prefix, "/example"),
		middleware.HandlerLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("another example"))
		})),
	)
}
