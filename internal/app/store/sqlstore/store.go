package sqlstore

import "database/sql"
import _ "github.com/lib/pq"

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
