package data

import (
	"context"
	"fmt"
	"github.com/witekblach/fridge/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"log/slog"
	"os"
)

type Ingredient struct {
	Name string
}

type CreateIngredientRequest struct {
	Name string `json:"name"`
}

func ShowAllIngredients() ([]Ingredient, error) {
	slog.Info("ShowAllIngredients called")
	coll := db.Mongo.Database("fridge").Collection("ingredients")

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var result []Ingredient
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	slog.Info(fmt.Sprintf("%+v", result))
	return result, nil
}

func AddIngredient(ingredient Ingredient) {
	slog.Info("AddIngredient called")

	coll := db.Mongo.Database("fridge").Collection("ingredients")

	_, err := coll.InsertOne(context.TODO(), ingredient)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
