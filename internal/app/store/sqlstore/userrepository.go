package sqlstore

import "github.com/Dennikoff/TodoAPI/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"insert into users values (default, $1, $2) returning id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}
