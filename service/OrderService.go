package service

import (
	"fmt"

	"github.com/eminoz/getir-backend-order-service/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	CreateNewOrder(order *model.Order, c *fiber.Ctx) (*mongo.InsertOneResult, error)
	GetOrders(userOrders primitive.D, c *fiber.Ctx) *mongo.SingleResult
	//UpdateOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) *mongo.SingleResult
	AddNewOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) (*mongo.UpdateResult, error)
	RemoveOneOrder(filter primitive.D, update primitive.D, c *fiber.Ctx) (*mongo.UpdateResult, error)
}
type OrderService struct {
	OrderRepo OrderRepository
}

func (repo OrderService) CreateNewOrder(c *fiber.Ctx) (*mongo.InsertOneResult, error) {
	order := new(model.Order)
	if err := c.BodyParser(&order); err != nil {
		return nil, c.Status(401).SendString(err.Error())
	}
	createdOrder, err := repo.OrderRepo.CreateNewOrder(order, c)
	if err != nil {
		return nil, c.Status(400).SendString(err.Error())
	}

	return createdOrder, err
}
func (repo OrderService) GetOrders(c *fiber.Ctx) (model.Order, error) {
	customerId := c.Params("id")
	_id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return model.Order{}, c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: _id}}
	order := repo.OrderRepo.GetOrders(filter, c)
	var result model.Order
	if err := order.Decode(&result); err != nil {
		return model.Order{}, c.Status(500).SendString("Something went wrong.")
	}

	return result, nil
}
func (repo OrderService) AddNewOrder(c *fiber.Ctx) (*mongo.UpdateResult, error) {
	userId := c.Params("id")
	_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}
	order := new(model.Order)

	if err := c.BodyParser(order); err != nil {
		fmt.Printf(err.Error())
		return nil, c.Status(401).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "product", Value: order.Product[0]}}}}

	updatedOrder, err := repo.OrderRepo.AddNewOrder(filter, update, c)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
func (repo OrderService) RemoveOneOrder(c *fiber.Ctx) (*mongo.UpdateResult, error) {
	removingOrder := new(model.RemoveOneOrder)

	if err := c.BodyParser(removingOrder); err != nil {
		fmt.Printf(err.Error())
		return nil, c.Status(401).SendString(err.Error())
	}

	_id, err := primitive.ObjectIDFromHex(removingOrder.OrderId)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}
	fmt.Print(removingOrder.OrderId, removingOrder.ProductId)
	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "product", Value: bson.D{
				{Key: "_id", Value: removingOrder.ProductId}}}}}}
	fmt.Println(update)
	updateOrder, err := repo.OrderRepo.RemoveOneOrder(filter, update, c)
	if err != nil {
		return nil, err
	}
	return updateOrder, nil
}
