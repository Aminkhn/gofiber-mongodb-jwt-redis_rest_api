package database

import (
	"context"
	"fmt"

	"github.com/aminkhn/mongo-rest-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// db is Mongo database object
var db *mongo.Database

// getting desired Collection from db
func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

// Establishing Mongo Connenction
func MongoConnect(config *config.Configuration) error {
	// Database settings (insert your own database name and connection URI)
	mongoURI := fmt.Sprintf("mongodb://%s/%s?connectTimeoutMS=5000", config.DBHost, config.DBName)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	dbname := fmt.Sprintf("%s", config.DBName)
	db = client.Database(dbname)
	return nil
}

// Closing database connctions
func CloseMongo() error {
	return db.Client().Disconnect(context.Background())
}
