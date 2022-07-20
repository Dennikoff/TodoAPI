package apiserver

import (
	"database/sql"
	"github.com/Dennikoff/TodoAPI/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(config Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	st := sqlstore.New(db)

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(st, sessionStore)
	srv.logger.Info("Starting api server")
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
