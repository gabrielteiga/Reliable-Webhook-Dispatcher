package webhooks

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/response"
)

func CreateWebhooks(w http.ResponseWriter, r *http.Request) {
	var webhooksRequest WebhooksRequest

	body := r.Body

	if err := json.NewDecoder(body).Decode(&webhooksRequest); err != nil {
		response.WriteJSON(
			w,
			http.StatusBadRequest,
			nil,
		)
	} else {
		response.WriteJSON(
			w,
			http.StatusOK,
			response.SuccessWithoutData(webhooksRequest.TargetUrl),
		)
	}
}
