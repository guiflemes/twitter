package controller

import (
	"encoding/json"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/users/service"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"github.com/guiflemes/twitter_clone/utils/http_utils"
	"net/http"
)

var (
	UserController userControllerInterface = &userController{}
)

type userController struct{}

type userControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		requestErr := errors.BadRequestErr("invalid json body")
		logger.Error("invalid json body", err)
		http_utils.ResponseError(w, *requestErr)
		return
	}

	result, createErr := service.UserService.Create(user)

	if createErr != nil {
		http_utils.ResponseError(w, *createErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)

}
