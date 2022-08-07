package users

import (
	"fmt"

	"github.com/yanuarultfah/bookstore_users-api/datasources/mysql/user_db"
	"github.com/yanuarultfah/bookstore_users-api/utils/date_utils"
	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *erorrs.RestErr {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return erorrs.NewNotFoundError(fmt.Sprintf("User %d Not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *erorrs.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return erorrs.NewBadRequestError(fmt.Sprintf("email %s already register", user.Email))
		}
		return erorrs.NewBadRequestError(fmt.Sprintf("User Already %d exist", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
