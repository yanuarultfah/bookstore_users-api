package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
	"github.com/yanuarultfah/bookstore_users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			message: "Invalid Json Body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status)
		return
	}

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
