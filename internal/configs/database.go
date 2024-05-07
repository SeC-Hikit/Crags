package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoDatabase struct {
	Client *mongo.Client
}

func NewMongoDatabase() MongoDatabase {
	uri := os.Getenv("MONGODB_URI")
	client, error := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))

	if error != nil {
		panic(error)
	}
	return MongoDatabase{Client: client}
}
