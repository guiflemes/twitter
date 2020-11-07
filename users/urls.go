package users

import (
	"github.com/guiflemes/twitter_clone/app/handlers"
	"github.com/guiflemes/twitter_clone/users/controller"
	"net/http"
)

func RegisterUrls() {
	userController := controller.UserController

	createUser := handlers.Route{
		Method:  http.MethodPost,
		Path:    "api/v1/users/",
		Handler: userController.Create,
	}

	handlers.RegisterRouter(createUser)

}
