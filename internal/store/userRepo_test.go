package store_test

import (
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := store.TestStore(t, dbUrl)
	defer teardown("users")

	u, err := s.GetUserRepo().Create(model.GetTestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, dbUrl)
	defer teardown("users")

	_, err := s.GetUserRepo().FindByEmail("test_find@test.com")

	assert.Error(t, err)

	ue, err := s.GetUserRepo().Create(model.GetTestUser(t))

	u, err := s.GetUserRepo().FindByEmail(ue.Email)

	assert.NoError(t, err)
	assert.Equal(t, ue, u)
}
