package main

import (
	"fmt"
	"net/http"
	"github.com/Underlike/go_api/controllers"
)

func main() {
	server()
}

func server() {
	fmt.Printf("Server start")
	fmt.Printf(controllers.GetValue())
	//http.HandleFunc("/", controllers.handlerRequest())
	//http.ListenAndServe(":8001", nil)
}