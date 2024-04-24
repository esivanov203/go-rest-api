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
