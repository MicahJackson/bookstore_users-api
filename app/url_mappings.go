package app

import (
	// "github.com/micahjackson/bookstore_users-api/controllers"
	"github.com/gin-gonic/gin"
)

func mapUrls() {
	// router.GET("/ping", controllers.Ping)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
