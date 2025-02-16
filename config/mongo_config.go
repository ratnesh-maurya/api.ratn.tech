package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB global variable
var MongoDB *mongo.Client

// ConnectMongoDB initializes a MongoDB connection
func ConnectMongoDB() {
	mongoURI := os.Getenv("MONGOURI")
	if mongoURI == "" {
		log.Fatal("MONGOURI environment variable not set")
	}

	// Creating a new MongoDB client and connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB successfully!")
	MongoDB = client
}

// GetCollection returns a collection from the given database
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("RatnTech").Collection(collectionName)
	return collection
}

func GetRepoCollection(collectionName string) *mongo.Collection {
	return GetCollection(MongoDB, collectionName)
}
