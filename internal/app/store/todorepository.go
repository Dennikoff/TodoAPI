package store

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindByUserID(id int) ([]*model.Todo, error)
}
