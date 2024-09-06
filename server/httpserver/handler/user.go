package handler

import (
	"encoding/json"
	"fmt"
	"github.com/salehborhani/todo-list/repository/mysqlrepo"
	"github.com/salehborhani/todo-list/server/httpserver/jwt"
	"github.com/salehborhani/todo-list/service/userservice"
	"io"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

		return
	}

	var userRequest userservice.RegisterRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return

	}

	err = json.Unmarshal(data, &userRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	mysqlRepo := mysqlrepo.New()
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Register(userRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{"message": "user created"}`))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

		return
	}

	var userLogin userservice.LoginRequest

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return

	}

	err = json.Unmarshal(data, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	mysqlRepo := mysqlrepo.New()
	userSvc := userservice.New(mysqlRepo)

	response, err := userSvc.Login(userLogin)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, response.Token)))
	if err != nil {
		return
	}

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

		return
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(401), http.StatusMethodNotAllowed)

		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := jwt.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(401), http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte(`You have logged in`))
}
