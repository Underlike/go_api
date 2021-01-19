package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Underlike/go_api/models"
	"github.com/Underlike/go_api/controllers"
	"github.com/gorilla/mux"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func main() {
	fmt.Printf(InfoColor, "Server start on http://localhost:8002")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", controllers.ArticlesHandler)
	router.HandleFunc("/categories", CategoriesHandler)
	http.ListenAndServe(":8002", router)
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(models.AllsCategories())
}