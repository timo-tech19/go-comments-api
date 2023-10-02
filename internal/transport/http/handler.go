package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// Custom Handler for managing routing and integration with comment service
type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

// Creates a new Handler instance with mounted routes, middleware and configured http Server.
func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	// inject middleware
	h.Router.Use(JSONMiddleware)
	h.Router.Use(Logger)
	h.Router.Use(TimeoutMiddleware)

	// configure server
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

// Connects handler functions to API endpoints(routes)
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})

	h.Router.HandleFunc("/api/v1/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
}

// Launches server.
// Server will shutdown gracefully incase of any errors.
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Shutting down gracefully...")
	return nil
}
