package store

import "github.com/WelchDragon/http-rest-api.git/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	GetUser(int) (*model.User, error)
	UpdateUser(int, *model.User) error
}
