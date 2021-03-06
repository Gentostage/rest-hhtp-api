package teststore

import (
	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
)

//UserRepository ....
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) UpdateUser(id int, user *model.User) error {
	for _, u := range r.users {
		if u.ID == id {
			r.users[id].About = user.About
			r.users[id].FullName = user.FullName
			r.users[id].Avatar = user.Avatar
			return nil
		}

	}
	return store.ErrRecordNorFound
}

func (r *UserRepository) GetUser(id int) (*model.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, store.ErrRecordNorFound
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	u.ID = len(r.users)
	return nil
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNorFound
	}
	return u, nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNorFound

}
