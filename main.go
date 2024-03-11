package main

import (
	"log"
	"net/http"

	handlers "github.com/ortizdavid/go-null-dbvalues/handlers"
)


func main() {
	
	mux := http.NewServeMux()

	var handler handlers.CustomerHandler

	mux.HandleFunc("GET /register", handler.RegisterForm)
	mux.HandleFunc("POST /register", handler.Register)

	log.Println("Listening to: http://127.0.0.1:8080")
	http.ListenAndServe(":8080", mux)
}