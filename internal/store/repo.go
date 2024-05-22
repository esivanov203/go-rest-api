package store

import "github.com/esivanov203/go-rest-api/internal/model"

type UserRepo interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}
