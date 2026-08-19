package webhooks

import (
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/httputils"
	"github.com/gabrielteiga/webhook-dispatcher/internal/api/middleware"
)

func RegisterRoutes(mux *http.ServeMux, prefix string) {
	mux.Handle(
		httputils.CreateRoute(http.MethodPost, prefix, "/webhooks"),
		middleware.HandlerLogger(http.HandlerFunc(CreateWebhooks)),
	)
}
