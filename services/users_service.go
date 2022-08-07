package services

import (
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

func CreateUser(user users.User) (*users.User, *erorrs.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
