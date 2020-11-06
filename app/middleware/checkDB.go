package middleware

import (
	"github.com/guiflemes/twitter_clone/app/config/data_base"
	"github.com/guiflemes/twitter_clone/logger"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r * http.Request) {
		_, err := data_base.CheckDBConnection()

		if err != nil{
			errMsg := "lost connection to the database"
			http.Error(w, errMsg, 500)
			logger.Error(errMsg, err)
			return
		}

		next.ServeHTTP(w, r)
	}
}