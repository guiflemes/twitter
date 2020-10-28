package service

import (
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
)

var (
	LoginService loginServiceInterface = &loginService{}
)

type loginService struct {}

type loginServiceInterface interface {
	Login(user domain.User) (*domain.User, *errors.RestErr)
}

func (l *loginService)Login(user domain.User) (*domain.User, *errors.RestErr){

	if err := user.Validate(); err != nil {
		return nil, err
	}

	password := user.Password

	if user.CheckUSerExist() {
		return nil, errors.BadRequestErr("User already exists")
	}

	if err := user.AuthorizedLogin(password); err != nil{
		return nil, err
	}

	return nil, nil

}
