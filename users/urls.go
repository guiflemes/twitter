package users

import (
	"fmt"
	"github.com/guiflemes/twitter_clone/app/handlers"
	"github.com/guiflemes/twitter_clone/users/controller"
)

var (
	endPoint = fmt.Sprintf("%s/users/", handlers.ApiVersion)
)

func init() {

	control := controller.UserController

	_ = handlers.Route{
		Method:  "POST",
		Path:    endPoint,
		Handler: control.Create,
	}

}
