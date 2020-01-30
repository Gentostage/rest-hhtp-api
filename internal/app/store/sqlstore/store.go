package sqlstore

import (
	"database/sql"

	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	db             *sql.DB
	userReposiroty *UserRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User ...
func (s *Store) User() store.UserRepository {
	if s.userReposiroty != nil {
		return s.userReposiroty
	}
	s.userReposiroty = &UserRepository{
		store: s,
	}
	return s.userReposiroty
}
