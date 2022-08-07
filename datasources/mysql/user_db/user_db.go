package userdb

import "database/sql"

var (
	usersDB *sql.DB
)

func init() {
	var err error
	usersDB, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
}
