package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/service"
	"github.com/yanuarultfah/bookstore_users-api/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		return
	}
	result, saveErr := service.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
