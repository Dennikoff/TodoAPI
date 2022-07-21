package sqlstore

import (
	"fmt"
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

func (r *TodoRepository) FindByUserID(id int) error {
	todo := &model.Todo{}
	rows, err := r.store.db.Query(
		"SELECT id, user_id, header, text, created_date FROM todo WHERE user_id = $1", id,
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.UserId, &todo.Header, &todo.Text, &todo.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(todo)
	}
	return nil
}
