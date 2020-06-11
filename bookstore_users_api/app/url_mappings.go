package app

import (
	"ea/bookstore_users_api/controllers/ping"
	"ea/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.POST("/users/search", users.SearchUser)
}
