package services

// services are for any and all business logic
// no connection point to http server

import (
	"github.com/micahjackson/bookstore_users-api/domain/dusers"
)

func CreateUser(user dusers.User) (*dusers.User, error) {
	return &user, nil
}
