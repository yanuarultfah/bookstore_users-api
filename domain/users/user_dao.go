package users

import (
	usersdb "github.com/yanuarultfah/bookstore_users-api/datasource/mysql/users_db"
	dateutils "github.com/yanuarultfah/bookstore_users-api/utils/date_utils"
	"github.com/yanuarultfah/bookstore_users-api/utils/errors"
	mysqlutils "github.com/yanuarultfah/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "insert into users(first_name,last_name,email,date_created) values(?,?,?,?);"
	queryGetUser    = "select id, first_name, last_name, email, date_created from users where id =?;"
	errorNoRows     = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = dateutils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	user.Id = userId

	return nil
}
