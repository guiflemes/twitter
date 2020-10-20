package domain

import (
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

func (u *User) Validate() *errors.RestErr {

	if u.Email = strings.TrimSpace(strings.ToLower(u.Email)); u.Email == "" {
		return errors.BadRequestErr("Invalid email")
	}

	if u.Password == ""{
		return errors.BadRequestErr("invalid password")
	}


	u.EncryptPassword()

	return nil

}

func (u *User) EncryptPassword(){
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hash)
}
