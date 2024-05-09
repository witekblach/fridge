package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Mongo *mongo.Client

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient() error {
	var (
		dbUser     = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
		dbPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("mongodb://%s:%s@%s:27017", dbUser, dbPassword, dbHost)
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	Mongo = client

	return nil
}
