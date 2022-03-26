package repository

import (
	"github.com/eminoz/getir-backend-order-service/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo OrderRepo) CreateNewOrder(order *model.Order, c *fiber.Ctx) (*mongo.InsertOneResult, error) {
	createdOrder, err := repo.collection.InsertOne(c.Context(), order)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}
	return createdOrder, nil
}
func (repo OrderRepo) GetOrders(userOrders primitive.D, c *fiber.Ctx) *mongo.SingleResult {
	respondedOrders := repo.collection.FindOne(c.Context(), userOrders)
	return respondedOrders
}
func (repo OrderRepo) AddNewOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) (*mongo.UpdateResult, error) {
	updatedOrder, err := repo.collection.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}
	return updatedOrder, nil

}
func (repo OrderRepo) RemoveOneOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) (*mongo.UpdateResult, error) {
	updatedOrder, err := repo.collection.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}
	return updatedOrder, nil
}

//update order
/*func (repo OrderRepo) UpdateOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) *mongo.SingleResult {
	updatedOrder := repo.collection.FindOneAndUpdate(c.Context(), filter, update)
	return updatedOrder
}*/
