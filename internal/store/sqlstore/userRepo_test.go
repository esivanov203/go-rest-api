package sqlstore_test

import (
	"github.com/esivanov203/go-rest-api/internal/model"
	"github.com/esivanov203/go-rest-api/internal/store"
	"github.com/esivanov203/go-rest-api/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, dbUrl)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.GetTestUser(t)

	err := s.GetUserRepo().Create(model.GetTestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, dbUrl)
	defer teardown("users")
	s := sqlstore.New(db)
	ue := model.GetTestUser(t)

	_, err := s.GetUserRepo().FindByEmail("test_find@test.com")
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	err = s.GetUserRepo().Create(ue)
	assert.NoError(t, err)

	u, err := s.GetUserRepo().FindByEmail(ue.Email)
	assert.NoError(t, err)
	u.Sanitize()
	assert.Equal(t, ue, u)
}

func TestUserRepo_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, dbUrl)
	defer teardown("users")
	s := sqlstore.New(db)
	ue := model.GetTestUser(t)

	_, err := s.GetUserRepo().Find(5555)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	err = s.GetUserRepo().Create(ue)
	assert.NoError(t, err)

	u, err := s.GetUserRepo().Find(ue.ID)
	assert.NoError(t, err)
	u.Sanitize()
	assert.Equal(t, ue, u)
}
