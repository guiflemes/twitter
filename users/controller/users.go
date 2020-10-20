package controller

import (
	"encoding/json"
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"net/http"
)

func Create(w http.ResponseWriter, r * http.Request){
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user); if err != nil{
		_ = errors.BadRequestErr("invalid Json body")
		return
	}


}