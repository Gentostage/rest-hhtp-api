package store_test

import (
	"testing"

	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	"github.com/stretchr/testify/assert"
)

//TestUserRepository_Create ...
func TestUserRepository_Create(t *testing.T) {
	s, teardowwn := store.TestStore(t, databaseURL)
	defer teardowwn("users")
	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardowwn := store.TestStore(t, databaseURL)
	defer teardowwn("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
