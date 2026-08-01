package middleware

import (
	"log"
	"net/http"
	"time"
)

func HandlerLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(
			"[IN] IP=%s - METHOD=%s - PATH=%s",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)

		log.Printf(
			"[OUT] IP=%s - METHOD=%s - PATH=%s - DURATION=%s",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
