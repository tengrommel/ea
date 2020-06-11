package users

import (
	"ea/bookstore_users_api/domain/users"
	"ea/bookstore_users_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var counter int

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO handle json error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
