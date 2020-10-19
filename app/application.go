package app

import (
	"github.com/guiflemes/twitter_clone/app/config/data_base"
	"github.com/guiflemes/twitter_clone/app/handlers"
	"github.com/guiflemes/twitter_clone/logger"
)

func StartApp() {
	status, err := data_base.CheckDBConnection()

	if status == 0 {
		logger.Error("No Connection to the database", err)
		return
	}

	handlers.Handlers()
}
