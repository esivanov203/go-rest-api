package store

import (
	"context"
	"github.com/esivanov203/go-rest-api/internal/model"
)

type UserRepo struct {
	store *Store
}

func NewUserRepo(store *Store) *UserRepo {
	return &UserRepo{store: store}
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	res, err := r.store.db.ExecContext(context.Background(), query, u.Email, u.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = int(id)
	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	query := "SELECT id, email, password, is_admin FROM users WHERE email = ?"
	if err := r.store.db.QueryRowContext(
		context.Background(), query, email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.IsAdmin); err != nil {
		return nil, err
	}
	return u, nil
}
