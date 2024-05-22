package sqlstore

import (
	"context"
	"database/sql"
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store"
)

type UserRepo struct {
	store.UserRepo
	store *Store
}

func NewUserRepo(store *Store) *UserRepo {
	return &UserRepo{store: store}
}

func (r *UserRepo) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.EncryptPwd(); err != nil {
		return err
	}

	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	res, err := r.store.db.ExecContext(context.Background(), query, u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)
	u.Sanitize()

	return nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	query := "SELECT id, email, password, is_admin FROM users WHERE email = ?"
	if err := r.store.db.QueryRowContext(
		context.Background(), query, email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.IsAdmin); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) Find(id int) (*model.User, error) {
	u := &model.User{}
	query := "SELECT id, email, is_admin FROM users WHERE id = ?"
	if err := r.store.db.QueryRowContext(
		context.Background(), query, id,
	).Scan(&u.ID, &u.Email, &u.IsAdmin); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
