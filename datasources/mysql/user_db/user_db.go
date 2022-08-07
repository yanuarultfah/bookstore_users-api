package user_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	datasourcename := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "", "127.0.0.1:3306", "users_db")
	var err error
	Client, err := sql.Open("mysql", datasourcename)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connected........")
}
