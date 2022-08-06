package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
	"github.com/yanuarultfah/bookstore_users-api/services"
	"github.com/yanuarultfah/bookstore_users-api/utils/erorrs"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := erorrs.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
