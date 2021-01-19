package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/Underlike/go_api/models"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsArticles())
}