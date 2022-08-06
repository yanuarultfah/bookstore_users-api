package users

import (
	"strings"

	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *erorrs.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return erorrs.NewBadRequestError("Invalid Json Body")
	}
}
