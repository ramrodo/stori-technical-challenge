package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Session      *mongo.Client
	Transactions *mongo.Collection
}

func (db MongoDB) ConnectDB() MongoDB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", dbUser, dbPassword)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return MongoDB{
		Session:      client,
		Transactions: client.Database("StoriDB").Collection("Transactions"),
	}
}

func (db MongoDB) CloseDB() {
	err := db.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
