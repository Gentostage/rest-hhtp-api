package model

import "testing"

//TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
		About:    "About ...",
		Avatar:   "/media/image.jpg",
		FullName: "User Users Useresevich",
	}
}
