package orderRoutes

import (
	"fiber-gorm/handlers/orderHandler"
	"github.com/gofiber/fiber/v2"
)

func Routes(orderRoute *fiber.App) {
	var order orderHandler.OrderHandler
	orderRoute.Post("/orders", order.CreateOrder)
	orderRoute.Get("/orders", order.GetAllOrders)
	orderRoute.Get("/orders/:id", order.GetOrder)

}
