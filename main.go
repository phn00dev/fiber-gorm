package main

import (
	"fiber-gorm/database"
	"fiber-gorm/routes/orderRoutes"
	"fiber-gorm/routes/productRoutes"
	"fiber-gorm/routes/userRoutes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	userRoutes.Routes(app)
	productRoutes.Routes(app)
	orderRoutes.Routes(app)
	_ = app.Listen(":3000")
}
