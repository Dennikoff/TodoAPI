package teststore

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type TodoRepository struct {
	store *Store
	todo  []*model.Todo
}

func (r *TodoRepository) Create(todo *model.Todo) error {

	r.todo = append(r.todo, todo)
	todo.ID = len(r.todo)
	return nil
}

func (r *TodoRepository) FindByUserID(id int) ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0, 2)
	for _, todo := range r.todo {
		if todo.UserId == id {
			todos = append(todos, todo)
		}
	}
	return todos, nil
}
