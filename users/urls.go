package users

import (
	"fmt"
	"github.com/guiflemes/twitter_clone/app/handlers"
)

var (
	endPoint = fmt.Sprintf("%s/users/", handlers.ApiVersion)
)

func init() {

	_ = handlers.Route{
		Method:  "GET",
		Path:    endPoint,
		Handler: nil,
	}

}
