package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create client for database
	client, err := mongo.NewClient(options.Client().ApplyURI(DatabaseURI()))
	if err != nil {
		log.Panicln("Unable to create new MongoDB client!!")
		log.Fatal(err)
	}
}