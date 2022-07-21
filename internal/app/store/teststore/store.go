package teststore

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
	todoRepository *TodoRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}
	return s.userRepository
}

func (s *Store) Todo() store.TodoRepository {
	if s.todoRepository == nil {
		s.todoRepository = &TodoRepository{
			store: s,
			todo:  make([]*model.Todo, 0),
		}
	}
	return s.todoRepository
}
