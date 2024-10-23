package users

import (
	"fmt"

	usersdb "github.com/yanuarultfah/bookstore_users-api/datasource/mysql/users_db"
	cryptoutils "github.com/yanuarultfah/bookstore_users-api/utils/crypto"
	dateutils "github.com/yanuarultfah/bookstore_users-api/utils/date_utils"
	"github.com/yanuarultfah/bookstore_users-api/utils/errors"
	mysqlutils "github.com/yanuarultfah/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "insert into users(first_name,last_name,email,date_created,status,password) values(?,?,?,?,?,?);"
	queryGetUser          = "select id, first_name, last_name, email, date_created,status from users where id =?;"
	errorNoRows           = "no rows in result set"
	queryUpdateUser       = "update users set first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser       = "delete from users where id=?;"
	queryFindUserByStatus = "select id,first_name,last_name,email,date_created,status from users where status = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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
	user.Status = "active"
	// user.Password = cryptoutils.GetMd5("Passwordt3sT")
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, cryptoutils.GetMd5(user.Password))
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

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()
	//select id,first_name,last_name,email,date_created,status
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}

	return results, nil

}
