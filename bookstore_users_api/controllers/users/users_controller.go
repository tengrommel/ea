package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var counter int

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
