package teststore

import (
	"github.com/esivanov203/go-rest-api/internal/store"
)

type Store struct {
	store.Store
	userRepo *UserRepo
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetUserRepo() store.UserRepo {
	if s.userRepo == nil {
		s.userRepo = NewUserRepo(s)
	}

	return s.userRepo
}
