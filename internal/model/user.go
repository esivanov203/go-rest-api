package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u *User) EncryptPwd() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}
