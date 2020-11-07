package domain

import (
	"context"
	"github.com/guiflemes/twitter_clone/app/config/data_base"
	"github.com/guiflemes/twitter_clone/logger"
	"github.com/guiflemes/twitter_clone/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	db         = data_base.Client.Database("twitter")
	collection = db.Collection("user")
)

func (user *User) CheckUserExist() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	query := bson.M{"email": user.Email}

	err := collection.FindOne(ctx, query).Decode(&user)

	if err != nil {
		return false
	}

	return true
}

func (user *User) Create() *errors.RestErr {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		logger.Error("error trying to save a new user", err)
		return errors.InternalServerError("database error")
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	user.ID = oid

	return nil
}

func (user *User) Get() *errors.RestErr {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	query := bson.M{"_id": user.ID}

	err := collection.FindOne(ctx, query).Decode(&user)

	if err != nil {
		logger.Error("error trying to get user by id", err)
		return errors.InternalServerError("database error")
	}

	return nil
}
