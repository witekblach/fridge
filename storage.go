package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=5432 sslmode=disable", dbname, dbuser, dbpassword, dbhost)
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		err = errors.New("error connecting to db")
		slog.Error(err.Error())

		return nil, err
	}

	return &Storage{db}, err
}
