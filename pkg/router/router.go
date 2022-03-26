package router

import (
	"github.com/eminoz/getir-backend-order-service/api"
	"github.com/eminoz/getir-backend-order-service/repository"
	"github.com/eminoz/getir-backend-order-service/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() *fiber.App {
	r := fiber.New()
	r.Use(cors.New())
	r.Use(logger.New())

	orderRepo := repository.DatabaseSetting()
	service := service.OrderService{OrderRepo: *orderRepo}
	controller := api.OrderController{OrderServices: service}

	post := r.Group("order")
	{
		post.Post("/", controller.CreateNewOrder)
		post.Get("/getorders/:id", controller.GetOrders)
		post.Put("/addneworder/:id", controller.AddNewOrder)
		post.Put("/removeoneorder/", controller.RemoveOneOrder)
	}
	return r
}
