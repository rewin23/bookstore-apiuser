 package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rewin23/bookstore-apiuser/domain/users"
	"github.com/rewin23/bookstore-apiuser/services"
	"github.com/rewin23/bookstore-apiuser/utils/errors"
)

func CreateUser(c *gin.Context) { 
	var user users.User
	if err := c.ShouldBindJSON(&user);  err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status: http.StatusBadRequest,
			Error:  "bad_request",
		}

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
        c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

