package main

import (
	"fmt"
	//"net/http"
	"controllers"
)

func main() {
	server()
}

func server() {
	fmt.Printf("Server start")
	fmt.Println(controllers.GetValue())
	//http.HandleFunc("/", controllers.handlerRequest())
	//http.ListenAndServe(":8001", nil)
}