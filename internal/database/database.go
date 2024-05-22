package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect initializes the connection to MongoDB
func Connect() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://kirillmitin34:HogEWwq7qFOr5eh8@crypto.idpoyke.mongodb.net/?retryWrites=true&w=majority&appName=crypto"))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}

	log.Println("Connected to MongoDB! 🚀")
}

// InsertTicker inserts a new ticker document into the specified collection
func InsertTicker(symbol string, lastPrice string, prevPrice24h string, timestamp time.Time) error {
	collection := client.Database("cryptoPricesDB").Collection("tickers")

	document := bson.M{
		"symbol":       symbol,
		"lastPrice":    lastPrice,
		"prevPrice24h": prevPrice24h,
		"time":         timestamp,
	}

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Printf("Failed to insert document into MongoDB: %v", err)
		return err
	}

	log.Printf("Inserted document into MongoDB: %v", document)
	return nil
}

// GetCollection returns a handle to a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("cryptoPricesDB").Collection(collectionName)
}
