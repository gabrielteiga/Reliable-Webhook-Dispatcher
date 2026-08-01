package api

import (
	"fmt"
	"net/http"
)

const SERVER_PORT string = ":8080"

func Run() {
	mux := http.NewServeMux()

	PopulateHandlers(mux)

	if err := http.ListenAndServe(SERVER_PORT, mux); err != nil {
		fmt.Println("shutdown http server")
	}
}
