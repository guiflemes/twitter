package users

import (
	"fmt"
	"github.com/guiflemes/twitter_clone/app/handlers"
	"github.com/guiflemes/twitter_clone/users/controller"
)

var (
	endPoint = fmt.Sprintf("%s/users/", handlers.ApiVersion)
)

func RegisterUrls() {
	control := controller.UserController

	createUser := handlers.Route{
		Method:  "POST",
		Path:    endPoint,
		Handler: control.Create,
	}

	handlers.RegisterRouter(createUser)

}
