package productHandler

import (
	"fiber-gorm/database"
	"fiber-gorm/models/productModel"
	"fiber-gorm/response/productResponse"
	"fiber-gorm/service/productService"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct{}

func (productHandler ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	var products []productModel.Products
	database.Database.Db.Find(&products)
	if len(products) == 0 {
		return c.Status(404).JSON("Products not found!!!")
	}
	responseProducts := []productResponse.ProductResponse{}
	for _, product := range products {
		responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func (productHandler ProductHandler) GetProduct(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	var product productModel.Products
	if err != nil {
		return c.Status(404).JSON("product not Found :id!=integer !!!")
	}
	err = productService.ProductService{}.FindProduct(productId, &product)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
	return c.Status(200).JSON(responseProduct)
}

func (productHandler ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product productModel.Products
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
	data := map[string]interface{}{
		"success": "Product created Successfully!!!",
		"Status":  201,
		"Product": responseProduct,
	}
	return c.Status(201).JSON(data)
}

func (productHandler ProductHandler) ProductUpdate(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("product not Found :id!=integer !!!")
	}
	var product productModel.Products
	err = productService.ProductService{}.FindProduct(productId, &product)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	var updateProduct productResponse.ProductUpdate

	if err = c.BodyParser(&updateProduct); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	product.ProductName = updateProduct.ProductName
	product.Price = updateProduct.Price
	product.TotalCount = updateProduct.TotalCount
	database.Database.Db.Save(&product)
	responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
	data := map[string]interface{}{
		"success": "Product updated Successfully!!!",
		"Status":  200,
		"Product": responseProduct,
	}
	return c.Status(201).JSON(data)
}

func (productHandler ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("product not Found :id!=integer !!!")
	}
	var product productModel.Products
	err = productService.ProductService{}.FindProduct(productId, &product)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.Database.Db.Unscoped().Delete(&product)
	data := map[string]interface{}{
		"success": "Product Deleted Successfully!!!",
		"Status":  200,
	}
	return c.Status(201).JSON(data)
}
