package repository

import (
	"github.com/eminoz/getir-backend-order-service/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func DatabaseSetting() *OrderRepo {
	database := database.GetDatabase()

	return &OrderRepo{db: database,
		collection: database.Collection("user")}
}
