package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/Underlike/go_api/models"
)

func main() {
	server()
}

func server() {
	fmt.Printf("Server start")
	http.HandleFunc("/", handlerRequest)
	http.ListenAndServe(":8001", nil)
}

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsArticles())
}