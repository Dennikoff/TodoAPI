package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Dennikoff/TodoAPI/internal/app/model"
	"github.com/Dennikoff/TodoAPI/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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
	s.router.Use(s.setLogger)

	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/create", s.handleUserCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/login", s.handleUserLogIn()).Methods(http.MethodPost)

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoAmI()).Methods(http.MethodGet)
	private.HandleFunc("/create", s.handleCreateTodo()).Methods(http.MethodPost)
	private.HandleFunc("/get", s.handleGetTodos()).Methods(http.MethodGet)
}

func (s *server) handleGetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := s.store.Todo().FindByUserID(r.Context().Value(ctxKeyUser).(*model.User).ID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, todos)
	}
}

func (s *server) handleCreateTodo() http.HandlerFunc {
	type request struct {
		Header string `json:"header"`
		Text   string `json:"text"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		todo := &model.Todo{
			Header: req.Header,
			Text:   req.Text,
		}
		todo.UserId = r.Context().Value(ctxKeyUser).(*model.User).ID
		if err := s.store.Todo().Create(todo); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, todo)
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyReqID, id)))
	})
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) setLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"bind_addr": r.RemoteAddr,
			"req_id":    r.Context().Value(ctxKeyReqID),
		})

		logger.Infof("Start %s %s\n", r.Method, r.RequestURI)
		start := time.Now()
		wr := &responseWriter{
			w,
			0,
		}
		next.ServeHTTP(wr, r)

		logger.Infof("Complete with %d, %s in %v", wr.code, http.StatusText(wr.code), time.Since(start))

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
