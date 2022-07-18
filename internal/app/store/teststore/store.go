package teststore

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type Store struct {
	userRepository *UserRepository
}

func (s *Store) Users() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}
	return s.userRepository
}

//TODO: add create and User func
