package model

import "testing"

func GetTestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "Pw!",
	}
}
