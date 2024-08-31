package entity

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Title        string
	Date         time.Time
	Status       bool
	Content      string
	UserID       uuid.UUID
	CategoryName string
}
