package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Initialize connection to MongoDB
func Connect() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Could not connect to MongoDB: ", err)
    }

    log.Println("Connected to MongoDB!")
}

// GetCollection returns a handle to a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("cryptoPricesDB").Collection(collectionName)
}
