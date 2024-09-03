package main

import (
	"fmt"
	"github.com/salehborhani/todo-list/server/httpserver/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/api/register/", handler.RegisterUser)

	// TODO - login user with the creds
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("Starting server on", server.Addr)
	if err := server.ListenAndServe(); err != nil {

	}
}
