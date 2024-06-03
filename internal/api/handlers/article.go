package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/durwankurnaik/glovs/internal/db"
	"github.com/durwankurnaik/glovs/internal/models"
	// "github.com/gorilla/mux"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	db.DB.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}