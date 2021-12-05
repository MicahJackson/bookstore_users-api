package services

// services are for any and all business logic
// no connection point to http server

import (
	"fmt"

	"github.com/micahjackson/bookstore_users-api/domain/dusers"
	"github.com/micahjackson/bookstore_users-api/utils/errors"
)

func CreateUser(user dusers.User) (*dusers.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	fmt.Println("successful save")
	return &user, nil
}
