package controller

import (
	"encoding/json"
	"github.com/guiflemes/twitter_clone/users/domain"
	"net/http"
)

func Create(w http.ResponseWriter, r * http.Request){
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user); if err != nil{

	}


}