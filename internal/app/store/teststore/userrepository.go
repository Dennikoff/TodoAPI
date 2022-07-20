package teststore

import (
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := r.users[email]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}
	return user, nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	for index := range r.users {
		if r.users[index].ID == id {
			return r.users[index], nil
		}
	}
	return nil, store.ErrorRecordNotFound
}
