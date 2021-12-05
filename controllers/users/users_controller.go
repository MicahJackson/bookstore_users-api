package users

// controller is for validating that we have all of the necessary parameters to handle the service request
// functions are the entry point of our application
// somehow we need an entry point, or http server, that handles all of the requests and redirects each request to a given controller based on the given "path"(?).

// application and controllers are the only layers that should interact with the server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micahjackson/bookstore_users-api/domain/dusers"
	"github.com/micahjackson/bookstore_users-api/services"
	"github.com/micahjackson/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) { // context in this case is the data behind -d {} that we're sending with the request
	var user dusers.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		// return bad request to the caller.
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
	c.String(http.StatusNotImplemented, "still not implemented")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
