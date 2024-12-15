package userRoutes

import (
	"fiber-gorm/handlers/userHandler"
	"github.com/gofiber/fiber/v2"
)

func Routes(userRoutes *fiber.App) {
	var userHandler userHandler.UserHandler
	userRoutes.Post("/users", userHandler.CreateUser)
	userRoutes.Get("/users", userHandler.GetAllUsers)
	userRoutes.Get("/users/:id", userHandler.GetUser)
	userRoutes.Put("/users/:id", userHandler.UpdateUser)
	userRoutes.Delete("/users/:id", userHandler.DeleteUser)
}
