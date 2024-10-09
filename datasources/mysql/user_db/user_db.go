package user_db

import (
	"database/sql"
	"log"

	// "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	// datasourcename := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8",
	// 	"root",
	// 	"localhost:3306",
	// 	"test",
	// )
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	var err error
	Client, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connected........")
}
