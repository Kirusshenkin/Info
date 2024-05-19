package auth

import (
	"context"
	"cryptoApi/internal/database"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExists(userID int) bool {
	collection := database.GetCollection("users")
	filter := bson.M{"id": userID}

	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

func AddUserToDatabase(dto CreateUserDTO) {
	collection := database.GetCollection("users")
	_, err := collection.InsertOne(context.Background(), dto)
	if err != nil {
		log.Printf("Error adding user to database: %v", err)
	}
}
