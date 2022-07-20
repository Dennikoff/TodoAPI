package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxKey uint8

var (
	errorUserNotAutheticate = errors.New("user is not authenticate")
)

const (
	SessionName        = "Authorized"
	ctxKeyUser  ctxKey = iota
	ctxKeyReqID ctxKey = iota
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, session sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: session,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)

	s.router.HandleFunc("/create", s.handleUserCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/login", s.handleUserLogIn()).Methods(http.MethodPost)
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyReqID, id)))
	})
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, SessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]

		if !ok {
			s.error(w, r, http.StatusUnauthorized, errorUserNotAutheticate)
			return
		}

		user, err := s.store.User().FindByID(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errorUserNotAutheticate)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))

	})
}

func (s *server) handleUserLogIn() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		user := &model.User{
			Email:             req.Email,
			EncryptedPassword: req.Password,
		}

		us, err := s.store.User().FindByEmail(user.Email)
		if err != nil || !us.ComparePassword(user.EncryptedPassword) {
			s.error(w, r, http.StatusUnauthorized, store.ErrorRecordNotFound)
			return
		}

		sessionStore, err := s.sessionStore.Get(r, SessionName)
		if err != nil {
			s.error(w, r, http.StatusInsufficientStorage, err)
			return
		}

		sessionStore.Values["user_id"] = us.ID

		if err := s.sessionStore.Save(r, w, sessionStore); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, us)

	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		user.Sanitize()
		s.respond(w, r, http.StatusCreated, user)

	}

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}
