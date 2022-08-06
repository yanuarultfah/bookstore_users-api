package services

import (
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

func CreateUser(User users.User) (*users.User, *erorrs.RestErr) {
	if err := users.Validate(); err != nil {

	}
	return nil, nil
}
