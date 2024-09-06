package main

import (
	"fmt"
	"github.com/salehborhani/todo-list/server/httpserver/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/user/register/", handler.RegisterUser)
	http.HandleFunc("/api/user/login/", handler.LoginUser)
	http.HandleFunc("/api/task/create/", handler.CreateTask)

	// TODO - login user with the creds
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("Starting server on", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
