package teststore

import "github.com/WelchDragon/http-rest-api.git/internal/app/store"

import "github.com/WelchDragon/http-rest-api.git/internal/app/model"

//Store ...
type Store struct {
	userReposiroty *UserRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//User ...
func (s *Store) User() store.UserRepository {
	if s.userReposiroty != nil {
		return s.userReposiroty
	}
	s.userReposiroty = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userReposiroty
}
