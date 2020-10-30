package service

import (
	"github.com/guiflemes/twitter_clone/app/auth"
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
)

var (
	SessionService sessionServiceInterface = &sessionService{}
)

type sessionService struct {}

type sessionServiceInterface interface {
	Login(user domain.User) (*auth.JwtToken, *errors.RestErr)
}

func (s *sessionService)Login(user domain.User) (*auth.JwtToken, *errors.RestErr){

	passwordAsString := user.Password

	if err := user.Validate(); err != nil {
		return nil, err
	}


	if ! user.CheckUSerExist() {
		return nil, errors.BadRequestErr("User doest not exists")
	}

	if err := user.AuthorizedLogin(passwordAsString); err != nil{
		return nil, err
	}

	tokenAsString, err := auth.CreateJWT(user)

	if err != nil{
		return nil, errors.BadRequestErr("error creating a token")
	}


	token := auth.JwtToken{
		TokenAsString: tokenAsString,
	}

	return &token, nil

}
