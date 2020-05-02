package user_controller

import "C"
import (
	"fmt"
	"github.com/Ramv4694/bookstore_users_api/domain/users"
	"github.com/Ramv4694/bookstore_users_api/services"
	"github.com/Ramv4694/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser( c * gin.Context){
	var user users.User
	if  err := c.ShouldBindJSON(&user); err != nil{
		//TODO HANDLE JSON ERROR
		restErr := errors.NewBadRequest("Invalid Json Body")
		c.JSON(restErr.Status,restErr)
		return


	}
	fmt.Println("user is ", user)
	result,saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Handle USER creation eror
		c.JSON(saveErr.Status,saveErr)
   		return
	}
	c.JSON(http.StatusCreated,result)

}


func GetUser(c * gin.Context){
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
    if userErr != nil{
    	err := errors.NewBadRequest("invalid user id")
		C.JSON(err.Status,err)
	}
	c.String(http.StatusNotImplemented,"mplemt mnee")
	result,getErr := services.GetUser(userId)
	if getErr != nil {
		//Handle USER creation eror
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusCreated,result)

}
/*
func SearchUser(c * gin.Context){

	c.String(http.StatusNotImplemented,"mplemt mnee")
}
*/
