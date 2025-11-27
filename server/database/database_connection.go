package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	MongoDb := os.Getenv("DATABASE_URI")
	if MongoDb == "" {
		log.Fatal("DATABASE_URI not found in environment variables")
	}

	fmt.Println("MongoDB URI:", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client = Connect()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("Database Name:", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}

	return collection
}
