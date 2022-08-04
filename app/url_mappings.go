package app

import "github.com/yanuarultfah/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
