package main

import (
	"github.com/joho/godotenv"
	"github.com/witekblach/fridge/db"
	"github.com/witekblach/fridge/router"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	err := actualMain()

	if err != nil {
		slog.Error(err.Error())
	}
}

func actualMain() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = db.NewMongoClient()
	if err != nil {
		return err
	}

	r := router.NewChiRouter()

	slog.Info("ListenAndServe on " + ":" + os.Getenv("APP_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
	if err != nil {
		return err
	}

	return nil
}
