package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	Id          string `json:"Id"`
	Date        string `json:"Date"`
	Transaction string `json:"Transaction"`
}

func (db MongoDB) InsertTransaction(tr Transaction) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, err := db.Transactions.InsertOne(ctx, tr)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	return objectId, nil
}

func (db MongoDB) GetAllTransactions() ([]Transaction, error) {
	var results []Transaction
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := db.Transactions.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem Transaction
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)
	return results, nil
}
