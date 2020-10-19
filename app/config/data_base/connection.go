package data_base

import (
	"context"
	"github.com/guiflemes/users_api/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	//Client Connection to the DATA-BASE
	Client *mongo.Client
)

func dbUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}

func init() {
	var err error

	clientOptions := options.Client().ApplyURI(dbUri())

	Client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Error("Error trying to stabilise MongoDb", err)
		panic(err)
	}

	if err = Client.Ping(context.TODO(), nil); err != nil {
		logger.Error("Pong err", err)
		panic(err)
	}

	logger.Info("Data-Base connection successfully")

}

func CheckDBConnection() (int, error) {

	err := Client.Ping(context.TODO(), nil)

	if err != nil {
		return 0, err
	}

	return 1, err
}
