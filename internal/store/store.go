package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config   *Config
	db       *sql.DB
	userRepo *UserRepo
}

func New(config *Config) *Store {
	return &Store{config: config}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.Url)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) GetUserRepo() *UserRepo {
	if s.userRepo == nil {
		s.userRepo = NewUserRepo(s)
	}

	return s.userRepo
}
