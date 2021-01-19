package controllers

import (
	"encoding/json"
	"net/http"
	"../models"
)

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsArticles())
}