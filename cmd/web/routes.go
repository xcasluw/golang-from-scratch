package main

import (
	"net/http"
	"xcasluw/golang-from-scratch/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	return mux
}
