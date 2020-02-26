package sqlstore_test

import (
	"testing"

	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestStore_UserRepository_UpdateUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	id := 0
	u1 := &model.User{}

	err := s.User().UpdateUser(id, u1)
	assert.Error(t, err)

	u2 := model.TestUser(t)
	s.User().Create(u2)

	u1.FullName = "Vasya Pupkin"
	u1.Avatar = "/link/image.jpg"
	u1.About = "About..."

	err = s.User().UpdateUser(u2.ID, u1)
	assert.NoError(t, err)

}

func TestStore_UserRepository_GetUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNorFound.Error())

	u := model.TestUser(t)
	u.Email = email

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}
