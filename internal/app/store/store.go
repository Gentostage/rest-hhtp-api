package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userReposiroty *UserRepository
}

//New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

//Close ...
func (s *Store) Close() {
	s.db.Close()
}

//User ...
func (s *Store) User() *UserRepository {
	if s.userReposiroty != nil {
		return s.userReposiroty
	}
	s.userReposiroty = &UserRepository{
		store: s,
	}
	return s.userReposiroty
}
