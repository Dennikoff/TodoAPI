package apiserver

import (
	"database/sql"
	"github.com/Dennikoff/TodoAPI/internal/app/store/sqlstore"
	"net/http"
)

func Start(config Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	st := sqlstore.New(db)

	srv := newServer(st)
	return http.ListenAndServe(config.Bind_addr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
