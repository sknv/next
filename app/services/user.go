package services

import (
	"net/http"

	"github.com/sknv/next/app/core/models"
)

type (
	User struct {
		WhoAmI *WhoAmI
	}

	UserResponse struct {
		User *models.User `json:"user"`
	}
)

func NewUser() *User {
	return &User{WhoAmI: NewWhoAmI()}
}

func (u *User) Me(r *http.Request) (*UserResponse, error) {
	currentUser, err := u.WhoAmI.GetCurrentUser(r)
	if err != nil {
		return nil, err
	}
	return &UserResponse{User: currentUser}, nil
}
