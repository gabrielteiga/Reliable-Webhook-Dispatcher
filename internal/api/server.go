package api

import (
	"fmt"
	"log"
	"net/http"
)

const SERVER_PORT string = ":8080"

func Run() {
	mux := http.NewServeMux()

	PopulateHandlers(mux)

	log.Printf("[INFO] - Starting server at port=%s", SERVER_PORT)

	if err := http.ListenAndServe(SERVER_PORT, mux); err != nil {
		fmt.Println("shutdown http server")
	}
}
