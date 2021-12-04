package users

// controller is for validating that we have all of the necessary parameters to handle the service request
// functions are the entry point of our application
// somehow we need an entry point, or http server, that handles all of the requests and redirects each request to a given controller based on the given "path"(?).

// application and controllers are the only layers that should interact with the server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micahjackson/bookstore_users-api/domain/dusers"
)

func CreateUser(c *gin.Context) {
	var user dusers.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "implemented.")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "still not implemented")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
