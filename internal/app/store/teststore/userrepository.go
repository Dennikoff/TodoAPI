package teststore

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {

}

