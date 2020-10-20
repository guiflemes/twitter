package domain

import (
	"context"
	"github.com/guiflemes/twitter_clone/app/config/data_base"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)


var (
	db = data_base.Client.Database("twitter")
	collection = db.Collection("user")
)

func (user *User) CheckUSerExist() bool{
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	query := bson.M{"email": user.Email}

	err := collection.FindOne(ctx, query)

	if err != nil{
		return false
	}

	return true
}

func (user *User) Create() *errors.RestErr{

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, user)

	if err != nil{

	}

	return nil
}
