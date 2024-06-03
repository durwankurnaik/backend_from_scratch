package router

import (
	"github.com/durwankurnaik/glovs/internal/api/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/articles", handlers.GetArticles).Methods("GET")

	return r
}