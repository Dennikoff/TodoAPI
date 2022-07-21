package sqlstore

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type TodoRepository struct {
	store *Store
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.store.db.QueryRow(
		"INSERT INTO todo (id, user_id, header, text, createdDate) VALUES (default, $1, $2, $3, default) RETURNING id",
		todo.UserId, todo.Header, todo.Text,
	).Scan(&todo.ID)
}
