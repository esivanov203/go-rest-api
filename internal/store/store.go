package store

type Store interface {
	GetUserRepo() UserRepo
}
