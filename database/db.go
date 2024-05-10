package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func ConnectDB() {
	// Initialize MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://atulranjan789:atul1234@cluster0.xr7e6vt.mongodb.net/")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

}

func GetDB() *mongo.Database {
	return client.Database("task_app")
}
