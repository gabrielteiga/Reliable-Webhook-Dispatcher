package health

import (
	"net/http"

	"github.com/gabrielteiga/webhook-dispatcher/internal/api/response"
)

// Status godoc
//
// @Summary 	Get Status
// @Description Get server status health
// @Tags		health
// @Produce		json
// @Success		200 {object} response.BaseResponse[any]
// @Router		/health [get]
func Status(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(
		w,
		http.StatusOK,
		response.SuccessWithoutData("It's everything fine!"),
	)
}
