package main

import (
	"net/http"
	"youtube-download/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define rotas separadas por método
	router.HandleFunc("/", handlers.POST).Methods(http.MethodPost)
	router.HandleFunc("/health", handlers.HandleHealthCheck).Methods(http.MethodGet)

	return router
}
