package app

import (
	"github.com/guiflemes/twitter_clone/session"
	"github.com/guiflemes/twitter_clone/users"
)

func init() {
	users.RegisterUrls()
	session.RegisterUrls()
}
