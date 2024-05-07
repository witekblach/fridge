package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
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
	slog.Info("AAAAAAAAAAAAAAAAA LOADED BITCH")
	storage, err := NewStorage()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result := storage.db.Exec(`
SELECT * 
FROM ingredients
`)
		fmt.Sprintf("%+v", result)
		w.Write([]byte("hello world"))
	})

	r.Post("/ingredient", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)

	return nil
}
