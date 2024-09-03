package userservice

import "github.com/salehborhani/todo-list/entity"

type Repository interface {
	Register(u entity.User) (entity.User, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
