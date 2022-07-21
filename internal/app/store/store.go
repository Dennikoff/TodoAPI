package store

type Store interface {
	User() UserRepository
	Todo() TodoRepository
}
