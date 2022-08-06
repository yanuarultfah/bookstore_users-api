package users

import (
	"errors"
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

func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, erorrs.NewBadRequestError("Invalid email address")
	}
}
