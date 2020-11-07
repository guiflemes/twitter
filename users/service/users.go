package service

import (
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	Create(domain.User) (*domain.User, *errors.RestErr)
}


func (u *userService) Create(user domain.User) (*domain.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if user.CheckUserExist() {
		return nil, errors.BadRequestErr("User already exists")
	}

	user.Status = domain.StatusState

	if err := user.Create(); err != nil {
		return nil, err
	}

	return &user, nil

}
