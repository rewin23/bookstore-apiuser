package services

import (
	"github.com/rewin23/bookstore-apiuser/domain/users"
	"github.com/rewin23/bookstore-apiuser/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}