package users

import (
	"fmt"
	"github.com/Ramv4694/bookstore_users_api/utils/errors"
)


var(
	usersDB = make(map[int64]*User)
)

func (user *User)Get()  * errors.RestErr{
	result := usersDB[user.Id]
	if result == nil{
		return errors.NewBadRequest(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}


func (user *User)Save() * errors.RestErr{
	current := usersDB[user.Id]
	if current != nil{
		if current.Email == user.Email{
			return errors.NewBadRequest(fmt.Sprintf("user Email %s already exists",user.Email))
		}

		return errors.NewBadRequest(fmt.Sprintf("user %d already exists",user.Id))

	}
	usersDB[user.Id] = user
	return nil

}

