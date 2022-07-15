package sqlstore

import "database/sql"

type Store struct {
	db *sql.DB
	// Repositories
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//func Repositories
