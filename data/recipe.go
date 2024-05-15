package data

import (
	"context"
	"github.com/witekblach/fridge/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"log/slog"
	"os"
)

type Ingredient struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name,omitempty" json:"name"`
	Amount string             `bson:"amount,omitempty" json:"amount"`
}

type CreateIngredientRequest struct {
	Name string `bson:"name" json:"name"`
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

func AddIngredient(ingredient Ingredient) error {
	coll := db.Mongo.Database("fridge").Collection("ingredients")

	_, err := coll.InsertOne(context.TODO(), ingredient)
	if err != nil {
		return err
	}

	return nil
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

	_, err := coll.DeleteOne(context.TODO(), bson.D{{"name", ingredient.Name}})
	if err != nil {
		return err
	}

	return nil
}
