package apiserver

import "errors"

var (
	ErrEmailOrPassword  = errors.New("Incorrect Email or Password")
	ErrNotAuthenticated = errors.New("You are not authenticated")
)
