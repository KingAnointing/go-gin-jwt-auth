package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create client for database
	client, err := mongo.NewClient(options.Client().ApplyURI(DatabaseURI()))
	if err != nil {
		log.Panicln("Unable to create new MongoDB client!!")
		log.Fatal(err)
	}

	//connect client to Database
	err = client.Connect(ctx)
	if err != nil {
		log.Panicln("Unable to connect client to database!!")
		log.Fatal(err)
	}

	// ping database to see if it is up
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panicln("Database server is down !!")
		log.Fatal(err)
	}

	// return client
	return client

}

var DbClient *mongo.Client = DatabaseConnection()

func Collections(collectionName string) *mongo.Collection {
	return DbClient.Database("cluster1").Collection(collectionName)
}
