package health

import (
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/httputils"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
)

func RegisterRoutes(mux *http.ServeMux, prefix string) {
	mux.Handle(
		httputils.CreateRoute(http.MethodGet, prefix, "/health"),
		middleware.HandlerLogger(http.HandlerFunc(Status)),
	)
}
