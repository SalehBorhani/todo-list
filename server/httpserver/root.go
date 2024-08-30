package main

import (
	"fmt"
	"github.com/salehborhani/todo-list/server/httpserver/request"
	"log"
	"net/http"
)

func main() {

	// TODO - register user with name, email and password
	http.HandleFunc("/api/register/", request.RegisterUser)
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("Starting server on", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	// TODO - login user with the creds

}
