package middleware

import (
	"net/http"
	"github.com/guiflemes/twitter_clone/app/auth"
)

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request){
		const BearerShema = "Bearer"
		tokenString := r.Header.Get("Authorization")
		jwtService := auth.JWTAuthService()
		token, err := jwtService.ValidateToken(tokenString)

		if !token.Valid{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}