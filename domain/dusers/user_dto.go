package dusers

// dto is the object(s) we are transferring from the controllers to the persistence layer/database
// this file defines the dto(s)
// the dto file SHOULD NOT CONTAIN ANY CALLS TO THE DATABASE(S)

import (
	"strings"

	"github.com/micahjackson/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
