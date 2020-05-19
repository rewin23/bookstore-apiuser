package users


import (
	"github.com/rewin23/bookstore-apiuser/utils/errors"
	"github.com/rewin23/bookstore-apiuser/utils/date_utils"

	"fmt"

) 

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User ) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d already exist", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}