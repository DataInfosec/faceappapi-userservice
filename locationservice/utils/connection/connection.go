package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	// "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection() *mongo.Collection {
	// Set client options
	// DB_URL := "mongodb+srv://elvis:0gbunike@cluster0-oqk2c.mongodb.net/test?retryWrites=true&w=majority"
	// clientOptions := options.Client().ApplyURI("mongodb+srv://"+os.Getenv("DB_URL")+"?retryWrites=true&w=majority")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_TABLE"))

	return collection
}
