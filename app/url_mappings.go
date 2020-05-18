package app

import (
	"github.com/rewin23/bookstore-apiuser/controllers/ping"
	"github.com/rewin23/bookstore-apiuser/controllers/users"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}