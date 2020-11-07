package controller

import (
	"encoding/json"
	"github.com/guiflemes/twitter_clone/app/auth"
	"github.com/guiflemes/twitter_clone/logger"
	service2 "github.com/guiflemes/twitter_clone/session/service"
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"github.com/guiflemes/twitter_clone/utils/http_utils"
	"net/http"
)

var(
	SessionController sessionControllerInterface = &sessionController{}
)

type sessionController struct {}

type sessionControllerInterface interface {
	Login(http.ResponseWriter, *http.Request)
}

func (s *sessionController)Login(w http.ResponseWriter, r *http.Request){

	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil{
		requestErr := errors.BadRequestErr("invalid json body")
		logger.Error("invalid json body", err)
		http_utils.ResponseError(w, *requestErr)
		return
	}

	isUserAuthenticated, errAuthenticated := service2.SessionService.Login(&user)

	if !isUserAuthenticated{
		http_utils.ResponseError(w, *errAuthenticated)
	}

	tokenAsString, tokenErr := auth.JWTAuthService().CreateJWT(user)

	if tokenErr != nil{
		http_utils.ResponseError(w, *errors.BadRequestErr("Not allowed"))
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, auth.JwtToken{AccessToken: tokenAsString})
}