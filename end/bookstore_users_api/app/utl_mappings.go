package app

import (
	"ea/end/bookstore_users_api/controllers/ping"
	"ea/end/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
}
