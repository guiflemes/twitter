package service

import (
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
)

var (
	SessionService sessionServiceInterface = &sessionService{}
)

type sessionService struct {}

type sessionServiceInterface interface {
	Login(user *domain.User) (bool, *errors.RestErr)
}

func (s *sessionService) Login(user *domain.User) (bool, *errors.RestErr){

	passwordAsString := user.Password

	if err := user.Validate(); err != nil {
		return false, err
	}

	if ! user.CheckUserExist() {
		return false, errors.BadRequestErr("User doest not exists")
	}

	if err := user.AuthorizedLogin(passwordAsString); err != nil{
		return false, err
	}

	return true, nil

}
