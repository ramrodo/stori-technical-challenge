package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Session      *mongo.Client
	Transactions *mongo.Collection
}

func (db MongoDB) ConnectDB() MongoDB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://ramrodo:wUSE9ID3cERGrJNS@clusterdb.6dulmuc.mongodb.net/?retryWrites=true&w=majority"))
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
