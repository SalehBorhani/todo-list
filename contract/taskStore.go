package contract

import "github.com/salehborhani/todo-list/entity"

type TaskWriteStore interface {
	Save(t entity.Task)
}

type TaskReadStore interface {
	Load() []entity.Task
}
