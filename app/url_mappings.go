package app

import (
	"github.com/yanuarultfah/bookstore_users-api/controllers/ping"
	"github.com/yanuarultfah/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
