package contract

import "github.com/salehborhani/todo-list/entity"

type CategoryWriteStore interface {
	Save(c entity.Category)
}

type CategoryReadStore interface {
	Load() []entity.Category
}
