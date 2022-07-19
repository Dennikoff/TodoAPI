package store

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
