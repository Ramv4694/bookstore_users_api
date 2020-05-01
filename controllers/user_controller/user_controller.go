package user_controller


import (
	"encoding/json"
	"fmt"
	"github.com/Ramv4694/bookstore_users_api/domain/users"
	"github.com/Ramv4694/bookstore_users_api/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser( c * gin.Context){
	var user users.User

	bytes,err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("The user is ", user.Id)
	if err != nil{
		//TODO: ERROR HANDLER

		return

	}
	if err := json.Unmarshal(bytes, &user); err != nil{

		//TODO: HANDLE JSON ERROR
		return

	}

	result,saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Handle USER creation eror
		return
	}
	fmt.Println(string(bytes))
	c.JSON(http.StatusCreated,result)

}


func GetUser(c * gin.Context){

	c.String(http.StatusNotImplemented,"mplemt mnee")

}
/*
func SearchUser(c * gin.Context){

	c.String(http.StatusNotImplemented,"mplemt mnee")
}
*/
