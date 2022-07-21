package sqlstore

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"log"
)

type TodoRepository struct {
	store *Store
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.store.db.QueryRow(
		"INSERT INTO todo (id, user_id, header, text, created_date) VALUES (default, $1, $2, $3, default) RETURNING id",
		todo.UserId, todo.Header, todo.Text,
	).Scan(&todo.ID)
}

func (r *TodoRepository) FindByUserID(id int) ([]*model.Todo, error) {
	rows, err := r.store.db.Query(
		"SELECT id, user_id, header, text, created_date FROM todo WHERE user_id = $1", id,
	)
	if err != nil {
		return nil, err
	}
	todos := make([]*model.Todo, 0, 4)
	for rows.Next() {
		todo := &model.Todo{}
		err := rows.Scan(&todo.ID, &todo.UserId, &todo.Header, &todo.Text, &todo.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
