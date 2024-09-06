package handler

import (
	"encoding/json"
	"fmt"
	"github.com/salehborhani/todo-list/repository/mysqlrepo"
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
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return

	}

	err = json.Unmarshal(data, &userRequest)
	if err != nil {
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	mysqlRepo := mysqlrepo.New()
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Register(userRequest)

	if err != nil {
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	_, err = w.Write([]byte(`{"message": "user created"}`))
	if err != nil {
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
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return

	}

	err = json.Unmarshal(data, &userLogin)
	if err != nil {
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	mysqlRepo := mysqlrepo.New()
	userSvc := userservice.New(mysqlRepo)

	res, err := userSvc.Login(userLogin)

	if err != nil {
		w.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, res.Token)))
	if err != nil {
		return
	}

}
