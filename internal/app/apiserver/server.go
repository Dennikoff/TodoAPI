package apiserver

import (
	"github.com/Dennikoff/TodoAPI/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	server *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	return &server{
		server: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
}
