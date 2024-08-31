package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/salehborhani/todo-list/entity"
	"github.com/salehborhani/todo-list/repository/filestorage"
	"github.com/salehborhani/todo-list/server/httpserver/request"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Create or Open fileStore and load to memoryStore `users` variable
	fs := filestorage.New("users.json")
	users := fs.Load()

	var reg request.UserRegister

	// Check if the Content-Type Header is application/json or not
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type. Expected application/json", http.StatusUnsupportedMediaType)

		return
	}
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	user := entity.User{
		ID:       uuid.New(),
		UserName: reg.UserName,
		Password: reg.Password,
	}

	users = append(users, user)
	fs.Save(user)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	// Create or Open fileStore and load to memoryStore `users` variable
	fs := filestorage.New("users.json")
	users := fs.Load()

	// Get the username and password from the basic auth header
	username, password, ok := r.BasicAuth()

	for _, user := range users {
		if user.UserName == username && user.Password == password && ok {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(`{"message": "You have logged-in!"}`)); err != nil {
				fmt.Println("Could not write message to stdout: ", err)
			}

			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	if _, err := w.Write([]byte(`User not found or invalid credential!`)); err != nil {
		fmt.Println("Could not write message to stdout: ", err)
	}

}
