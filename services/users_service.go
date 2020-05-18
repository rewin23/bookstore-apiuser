package services

import (
	"github.com/rewin23/bookstore-apiuser/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}