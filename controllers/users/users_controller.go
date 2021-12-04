package users

// controller is for validating that we have all of the necessary parameters to handle the service request
// functions are the entry point of our application
// somehow we need an entry point, or http server, that handles all of the requests and redirects each request to a given controller based on the given "path"(?).

// application and controllers are the only layers that should interact with the server

import (
	"fmt"
	"net/http"

	"example.go/github.com/micahjackson/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/micahjackson/bookstore_users-api/domain/dusers"
	"github.com/micahjackson/bookstore_users-api/services"
)

func CreateUser(c *gin.Context) {
	var user dusers.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		//TODO: return bad request to the caller.
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation errorf
		return
	}
	c.JSON(http.StatusCreated, result)
	fmt.Println("works")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "still not implemented")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
