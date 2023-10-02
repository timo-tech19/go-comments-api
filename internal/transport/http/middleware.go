package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// sets http response header to application/json
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}

// Logs request metadata to the console
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("handled request")

		next.ServeHTTP(w, r)
	})
}

// sets a timeout for processing each request
func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
