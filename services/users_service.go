package services

import (
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

func GetUser(userId int64) (*users.User, *erorrs.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *erorrs.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
