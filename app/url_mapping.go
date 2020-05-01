package app

import(
	"github.com/Ramv4694/bookstore_users_api/controllers/ping_controller"
	"github.com/Ramv4694/bookstore_users_api/controllers/user_controller"
)

func mapUrls(){

	router.GET("/ping", ping_controller.Ping)
	router.POST("/users",user_controller.CreateUser)
	router.GET("/users/:user_id",user_controller.GetUser)

}
