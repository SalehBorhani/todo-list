package memorystorage

import "github.com/salehborhani/todo-list/entity"

type UserStorage struct {
	users []entity.User
}

func (us UserStorage) NewUserStorage() *UserStorage {
	return &UserStorage{
		users: make([]entity.User, 0),
	}
}
