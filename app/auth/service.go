package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/guiflemes/twitter_clone/users/domain"
	"time"
)


func CreateJWT(user domain.User) (string, error){

	payload := jwt.MapClaims{
		"_id": user.ID.Hex(),
		"email": user.Email,
		"first_name": user.FirstName,
		"last_name": user.LastName,
		"birth_day": user.BirthDay,
		"bibliography": user.Bibliography,
		"web_site": user.WebSite,
		"location": user.Location,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenAsString, err := token.SignedString([]byte("twitter_clone.io"))


	if err != nil{
		logger.Error("error to signed jwt token to string", err)
		return tokenAsString, err
	}

	return tokenAsString, nil
}