package request

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/salehborhani/todo-list/entity"
	"github.com/salehborhani/todo-list/repository/filestorage"
	"net/http"
)

type Register struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// Create or Open fileStore and load to memoryStore `users` variable
	fs := filestorage.New("users.json")
	users := fs.Load()

	var reg Register
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
