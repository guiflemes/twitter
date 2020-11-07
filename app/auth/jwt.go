package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/guiflemes/twitter_clone/users/domain"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"
)

type JwtService interface {
	CreateJWT(domain.User) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}

type claim struct {
	Password string             `json:"password"`
	ID       primitive.ObjectID `bson:"_id" json:"id, omitempty"`
	jwt.StandardClaims
}

type JwtToken struct {
	AccessToken string `json:"access_token"`
}

type jwtService struct {
	secretKey string
	issure    string
}

func JWTAuthService() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SECRET_KEY")

}

func (service *jwtService) CreateJWT(user domain.User) (string, error) {

	payload := jwt.MapClaims{
		"_id":          user.ID.Hex(),
		"email":        user.Email,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"birth_day":    user.BirthDay,
		"bibliography": user.Bibliography,
		"web_site":     user.WebSite,
		"location":     user.Location,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenAsString, err := token.SignedString([]byte("twitter_clone.io"))

	if err != nil {
		logger.Error("error to signed jwt token to string", err)
		return tokenAsString, err
	}

	return tokenAsString, nil
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		encodedToken, &claim{}, func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(service.secretKey), nil
		})
}
