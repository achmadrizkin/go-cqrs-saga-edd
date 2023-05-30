package mongodb

import (
	"context"
	"errors"
	"go-cqrs-saga-edd/order-query/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbConn(context context.Context) (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://root:root@clusterpractice.bhieixl.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context, opts)
	if err != nil {
		return nil, errors.New("connection error (mongodb): " + err.Error())
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(config.Config("MONGO_DB_NAME")).RunCommand(context, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, errors.New("ping error to mongo db: " + err.Error())
	}

	return client, nil
}

func MongoCollection(col string, client *mongo.Client) *mongo.Collection {
	return client.Database(config.Config("MONGO_DB_NAME")).Collection(col)
}
