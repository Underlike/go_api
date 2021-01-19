package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/Underlike/go_api/controllers"
	"github.com/gorilla/mux"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func main() {
	initializeRouter()
}

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/articles", controllers.ArticlesHandler)
	router.HandleFunc("/categories", controllers.CategoriesHandler)

	srv := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:5002",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	fmt.Printf(InfoColor, "Server start on http://localhost:5002")
	log.Fatal(srv.ListenAndServe())
}