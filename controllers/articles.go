package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/Underlike/go_api/models"
)

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsArticles())
}

func getValue() string {
	return "'OK'"
}