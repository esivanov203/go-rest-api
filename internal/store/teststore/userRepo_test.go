package teststore_test

import (
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store"
	"github.com/esivanov203/go-rest-api/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s := teststore.New()
	u := model.GetTestUser(t)

	err := s.GetUserRepo().Create(model.GetTestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByID(t *testing.T) {
	s := teststore.New()
	ue := model.GetTestUser(t)

	_, err := s.GetUserRepo().FindByEmail("test_find@test.com")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	err = s.GetUserRepo().Create(ue)
	assert.NoError(t, err)

	u, err := s.GetUserRepo().FindByEmail(ue.Email)

	assert.NoError(t, err)
	assert.Equal(t, ue, u)
}
