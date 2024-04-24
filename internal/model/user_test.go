package model_test

import (
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_EncryptPwd(t *testing.T) {
	u := model.GetTestUser(t)
	pwd := u.Password
	err := u.EncryptPwd()

	assert.NoError(t, err)
	assert.NotEqual(t, pwd, u.Password)
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.GetTestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty_email",
			user: func() *model.User {
				u := model.GetTestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short_email",
			user: func() *model.User {
				u := model.GetTestUser(t)
				u.Email = "rrrr.ru"
				return u
			},
			isValid: false,
		},
		{
			name: "empty_password",
			user: func() *model.User {
				u := model.GetTestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short_password",
			user: func() *model.User {
				u := model.GetTestUser(t)
				u.Password = "t5"
				return u
			},
			isValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.NoError(t, test.user().Validate())
			} else {
				assert.Error(t, test.user().Validate())
			}
		})
	}

	u := model.GetTestUser(t)
	assert.NoError(t, u.Validate())
}
