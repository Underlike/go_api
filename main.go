package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Underlike/go_api/models"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func main() {
	fmt.Printf(InfoColor, "Server start on http://localhost:8001")
	http.HandleFunc("/", handlerRequest)
	http.ListenAndServe(":8001", nil)
}

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsArticles())
}