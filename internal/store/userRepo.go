package store

import "github.com/esivanov203/go-rest-api/internal/model"

type UserRepo struct {
	store *Store
}

func NewUserRepo(store *Store) *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	res, err := r.store.db.Exec(query, u.Email, u.Password)
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
	return nil, nil
}
