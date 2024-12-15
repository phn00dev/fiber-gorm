package productRoutes

import (
	"fiber-gorm/handlers/productHandler"
	"github.com/gofiber/fiber/v2"
)

func Routes(productRoutes *fiber.App) {
	var product productHandler.ProductHandler
	productRoutes.Post("/products", product.CreateProduct)
	productRoutes.Get("/products", product.GetAllProducts)
	productRoutes.Get("/products/:id", product.GetProduct)
	productRoutes.Put("/products/:id", product.ProductUpdate)
	productRoutes.Delete("/products/:id", product.DeleteProduct)
}
