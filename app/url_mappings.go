package app

import (
	"github.com/yanuarultfah/bookstore_users-api/controllers/ping"
	"github.com/yanuarultfah/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
	router.POST("/users/:user_id", users.Delete)
	router.GET("/internal/users/search/", users.Search)
}
