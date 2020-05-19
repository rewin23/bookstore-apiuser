package services

import (
	"github.com/rewin23/bookstore-apiuser/domain/users"
	"github.com/rewin23/bookstore-apiuser/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err:= user.Save(); err !=nil {
		return nil, err
	}
	return &user , nil
}