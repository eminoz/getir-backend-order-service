package database

import (
	"context"
	"time"

	"github.com/eminoz/getir-backend-order-service/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func Setup() error {

	config := config.GetConfig()
	var database *mongo.Database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDb))

	if err != nil {
		panic(err)
	}

	database = client.Database("user")

	Database = database
	return nil
}

func GetDatabase() *mongo.Database {
	return Database
}
