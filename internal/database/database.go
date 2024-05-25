package database

import (
	"context"

	"event_manager/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Mongo *mongo.Database
}

var db Database

func Connect() error {
	opts := options.Client().ApplyURI(config.Cfg().Mongo.URL)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db.Mongo = client.Database(config.Cfg().Mongo.DB)

	return nil
}

func Db() *Database {
	return &db
}
