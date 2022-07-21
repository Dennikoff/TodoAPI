package sqlstore

import (
	"database/sql"
	"github.com/Dennikoff/TodoAPI/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	todoRepository *TodoRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Todo() store.TodoRepository {
	if s.todoRepository == nil {
		s.todoRepository = &TodoRepository{
			store: s,
		}
	}
	return s.todoRepository
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}

//func Repositories
