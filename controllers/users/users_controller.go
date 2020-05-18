 package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rewin23/bookstore-apiuser/domain/users"
	"github.com/rewin23/bookstore-apiuser/services"
	"fmt"
)

func CreateUser(c *gin.Context) { 
	var user users.User
	if err := c.ShouldBindJSON(&user);  err != nil {
		fmt.Println(err)
		// TODO: return a bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creaton error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

