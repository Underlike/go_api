package main

import (
	"fmt"
	"net/http"
)

func main() {
	server()
}

func server() {
	fmt.Printf("Server start")
	http.HandleFunc("/", controllers.handlerRequest())
	http.ListenAndServe(":8001", nil)
}