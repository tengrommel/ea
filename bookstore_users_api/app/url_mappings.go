package app

import "ea/bookstore_users_api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", controllers.CreateUser)
	router.POST("/users/search", controllers.SearchUser)
}
