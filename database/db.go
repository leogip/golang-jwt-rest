package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/leogip/golang-jwt-rest/logger"
)

var (
	log = logger.New("Database")
	databaseName string = "godb"
	// CollectionUser ...
	CollectionUser *mongo.Collection
	// CollectionTokens ...
	CollectionTokens *mongo.Collection
)

// Connect to db
func Connect(databaseURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.WithError(err).Errorf("Can't open database")
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.WithError(err).Errorf("Can't open database")
		return err
	}

	CollectionUser = getCollectionUser(client)
	CollectionTokens = getCollectionTokens(client)

	return nil
}

func getCollectionUser(c *mongo.Client) *mongo.Collection {
	collection := c.Database(databaseName).Collection("users")
	return collection
}

func getCollectionTokens(c *mongo.Client) *mongo.Collection {
	collection := c.Database(databaseName).Collection("tokens")
	return collection
}
