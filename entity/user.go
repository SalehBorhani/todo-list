package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
}
