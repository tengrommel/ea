package app

import "ea/end/bookstore_users_api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
