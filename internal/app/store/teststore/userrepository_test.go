package teststore_test

import (
	"testing"

	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestStore_UserRepository_UserUpdate(t *testing.T) {
	s := teststore.New()

	id := 0
	u1 := &model.User{}

	err := s.User().UpdateUser(id, u1)
	assert.EqualError(t, err, store.ErrRecordNorFound.Error())

	u2 := model.TestUser(t)
	s.User().Create(u2)

	err = s.User().UpdateUser(u2.ID, u2)
	assert.NoError(t, err)

}

func TestStore_UserRepository_GetUser(t *testing.T) {

	s := teststore.New()

	id := 0

	u, err := s.User().GetUser(id)
	assert.EqualError(t, err, store.ErrRecordNorFound.Error())

	u1 := model.TestUser(t)
	s.User().Create(u1)

	u, err = s.User().GetUser(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

//TestUserRepository_Create ...
func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNorFound.Error())

	u := model.TestUser(t)
	u.Email = email

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}
