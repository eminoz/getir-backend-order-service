package api

import (
	"github.com/eminoz/getir-backend-order-service/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	CreateNewOrder(c *fiber.Ctx) (*mongo.InsertOneResult, error)
	GetOrders(c *fiber.Ctx) (model.Order, error)
	AddNewOrder(c *fiber.Ctx) (*mongo.UpdateResult, error)
	RemoveOneOrder(c *fiber.Ctx) (*mongo.UpdateResult, error)
}
type OrderController struct {
	OrderServices OrderService
}

func (o OrderController) CreateNewOrder(c *fiber.Ctx) error {
	createdOrder, err := o.OrderServices.CreateNewOrder(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(201).JSON(createdOrder)
}
func (o OrderController) GetOrders(c *fiber.Ctx) error {
	userOrders, err := o.OrderServices.GetOrders(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(userOrders)
}
func (o OrderController) AddNewOrder(c *fiber.Ctx) error {
	updatedOrder, err := o.OrderServices.AddNewOrder(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusCreated).JSON(updatedOrder)
}
func (o OrderController) RemoveOneOrder(c *fiber.Ctx) error {
	updatedOrder, err := o.OrderServices.RemoveOneOrder(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusCreated).JSON(updatedOrder)

}
