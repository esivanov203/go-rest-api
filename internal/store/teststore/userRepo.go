package teststore

import (
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store"
)

type UserRepo struct {
	store.UserRepo
	store *Store
	users map[string]*model.User
}

func NewUserRepo(store *Store) *UserRepo {
	return &UserRepo{
		store: store,
		users: make(map[string]*model.User),
	}
}

func (r *UserRepo) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.EncryptPwd(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
