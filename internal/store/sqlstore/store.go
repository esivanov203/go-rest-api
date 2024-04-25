package sqlstore

import (
	"database/sql"
	"github.com/esivanov203/go-rest-api/internal/store"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	store.Store
	db       *sql.DB
	userRepo *UserRepo
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserRepo() store.UserRepo {
	if s.userRepo == nil {
		s.userRepo = NewUserRepo(s)
	}

	return s.userRepo
}
