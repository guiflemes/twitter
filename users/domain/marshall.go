package domain

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PublicUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"first_name" json:"first_name,omitempty"`
	LastName     string             `bson:"last_name" json:"last_name,omitempty"`
	BirthDay     time.Time          `bson:"birth_day" json:"birth_day,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Bibliography string             `bson:"bibliography" json:"bibliography,omitempty"`
	Location     string             `bson:"location" json:"location,omitempty"`
	WebSite      string             `bson:"web_site" json:"web_site,omitempty"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    string             `bson:"created_at" json:"created_at"`
}

type PrivateUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"first_name" json:"first_name,omitempty"`
	LastName     string             `bson:"last_name" json:"last_name,omitempty"`
	BirthDay     time.Time          `bson:"birth_day" json:"birth_day,omitempty"`
	Email        string             `bson:"email" json:"email,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Bibliography string             `bson:"bibliography" json:"bibliography,omitempty"`
	Location     string             `bson:"location" json:"location,omitempty"`
	WebSite      string             `bson:"web_site" json:"web_site,omitempty"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    string             `bson:"created_at" json:"created_at"`
}


func (user *User) Marshal(isPublic bool) interface{}{
	var publicUser PublicUser
	var privateUser PrivateUser

	marshal := func(user *User) interface {}{
		userJson, _ := json.Marshal(user)

		if isPublic{
			if err := json.Unmarshal(userJson, &publicUser); err != nil{
				return nil
			}

			return publicUser
		}

		if err := json.Unmarshal(userJson, &privateUser);err != nil{
			return nil
		}

		return privateUser
	}

	return marshal(user)
}