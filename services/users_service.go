package services

import "github.com/yanuarultfah/bookstore_users-api/domain/users"

func CreateUser(User users.User) (*users.User, error) {
	return &User, nil
}
