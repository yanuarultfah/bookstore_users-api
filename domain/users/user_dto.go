package users

import (
	"strings"

	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *erorrs.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return erorrs.NewBadRequestError("Invalid Email")
	}
	return nil
}
