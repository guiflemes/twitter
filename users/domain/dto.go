package domain

import (
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const (
	StatusState = "active"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"first_name" json:"first_name,omitempty"`
	LastName     string             `bson:"last_name" json:"last_name,omitempty"`
	BirthDay     time.Time          `bson:"birth_day" json:"birth_day,omitempty"`
	Email        string             `bson:"email" json:"email,omitempty"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Bibliography string             `bson:"bibliography" json:"bibliography,omitempty"`
	Location     string             `bson:"location" json:"location,omitempty"`
	WebSite      string             `bson:"web_site" json:"web_site,omitempty"`
	Status       string             `bson:"status" json:"status"`
}

func (user *User) Validate() *errors.RestErr {

	if user.Email = strings.TrimSpace(strings.ToLower(user.Email)); user.Email == "" {
		return errors.BadRequestErr("Invalid email")
	}

	if user.Password == ""{
		return errors.BadRequestErr("invalid password")
	}


	user.EncryptPassword()

	return nil

}

func (user *User) EncryptPassword(){
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(hash)
}


func (user *User) AuthorizedLogin(password string) *errors.RestErr{

	inputPassword := []byte(password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), inputPassword)

	if err != nil{
		logger.Error("password error", err)
		return errors.ForbiddenError("invalid credentials to email or password")
	}

	return nil
}