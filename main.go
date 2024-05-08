package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	err := actualMain()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func actualMain() error {
	godotenv.Load()
	slog.Info("NewMongoClient")
	client, err := NewMongoClient()
	if err != nil {
		return err
	}

	slog.Info("NewRouter")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", client.ShowAllIngredients)

	r.Post("/ingredient", client.AddIngredient)

	slog.Info("ListenAndServe on " + ":" + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)

	return nil
}

func (mc MongoClient) AddIngredient(w http.ResponseWriter, r *http.Request) {
	slog.Info("AddIngredient called")

	coll := mc.client.Database("fridge").Collection("ingredients")

	catFact := bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}

	slog.Info("InsertOne")
	_, err := coll.InsertOne(context.TODO(), catFact)
	if err != nil {
		slog.Error("InsertOne")
	}
}

func (mc MongoClient) ShowAllIngredients(w http.ResponseWriter, r *http.Request) {
	slog.Info("ShowAllIngredients called")
	coll := mc.client.Database("fridge").Collection("ingredients")

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		slog.Error(err.Error())
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

type MongoClient struct {
	client *mongo.Client
}

func (mc MongoClient) start() error {
	return nil
}

func NewMongoClient() (*MongoClient, error) {
	var (
		dbUser     = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
		dbPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("mongodb://%s:%s@%s:27017", dbUser, dbPassword, dbHost)
	)

	slog.Info(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return &MongoClient{client}, err
}
