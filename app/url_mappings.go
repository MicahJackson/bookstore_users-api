package app

import (
	"github.com/micahjackson/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
