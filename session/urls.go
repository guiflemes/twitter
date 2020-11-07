package session

import (
	"github.com/guiflemes/twitter_clone/app/handlers"
	"github.com/guiflemes/twitter_clone/session/controller"
	"net/http"
)

func RegisterUrls() {

	sessionController := controller.SessionController

	login := handlers.Route{
		Method:  http.MethodPost,
		Path:    "api/v1/login/",
		Handler: sessionController.Login,
	}

	handlers.RegisterRouter(login)

}
