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
	Name string `json:"name"`
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
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Fatal(err)
	}

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

func GetIngredient(name string) (*Ingredient, error) {
	coll := db.Mongo.Database("fridge").Collection("ingredients")

	cursor, err := coll.Find(context.TODO(), bson.D{{"name", name}})
	if err != nil {
		return nil, err

	}

	var result Ingredient
	cursor.Next(context.TODO())
	if err = cursor.Decode(&result); err != nil {
		log.Fatal(err)
	}

	return &result, nil
}

func RemoveIngredient(ingredient Ingredient) error {
	coll := db.Mongo.Database("fridge").Collection("ingredients")

	slog.Info(ingredient.Name)

	deletedCount, err := coll.DeleteOne(context.TODO(), bson.D{{"name", ingredient.Name}})
	if err != nil {
		return err
	}

	slog.Info(fmt.Sprint(deletedCount.DeletedCount))

	return nil
}
