package sqlstore

import (
	"database/sql"
	"github.com/WelchDragon/http-rest-api.git/internal/app/model"
	"github.com/WelchDragon/http-rest-api.git/internal/app/store"
)

//UserRepository ..
type UserRepository struct {
	store *Store
}

//UpdateUser ...
func (r *UserRepository) UpdateUser(id int, u *model.User) error {
	if result, err := r.store.db.Exec("UPDATE users SET about=$2, avatar=$3, full_name= $4 where id = $1",
		id,
		u.About,
		u.Avatar,
		u.FullName,
	); err != nil {
		return err
	} else {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return store.ErrRecordNorFound
		}
	}
	return nil
}

//GetUser ...
func (r *UserRepository) GetUser(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Avatar,
		&u.FullName,
		&u.About,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNorFound
		}
	}
	return u, nil
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id=$1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNorFound
		}
		return nil, err
	}
	return u, nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNorFound
		}
		return nil, err
	}
	return u, nil
}
