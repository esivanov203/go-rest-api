package apiserver

import (
	"database/sql"
	"github.com/esivanov203/go-rest-api/internal/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDb(config.DbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDb(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
