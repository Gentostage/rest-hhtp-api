package sqlstore_test

import (
	"testing"

	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

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
