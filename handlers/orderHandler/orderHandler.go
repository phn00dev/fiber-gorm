package orderHandler

import (
	"fiber-gorm/database"
	"fiber-gorm/models/orderModel"
	"fiber-gorm/models/productModel"
	"fiber-gorm/models/userModel"
	"fiber-gorm/response/orderResponse"
	"fiber-gorm/response/productResponse"
	"fiber-gorm/response/userResponse"
	"fiber-gorm/service/orderService"
	"fiber-gorm/service/productService"
	"fiber-gorm/service/userService"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct{}

func (orderHandler OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	var orders []orderModel.Orders
	database.Database.Db.Find(&orders)
	if len(orders) == 0 {
		return c.Status(404).JSON("orders not found !!! orders table empty")
	}

	responseOrders := []orderResponse.OrderResponse{}

	for _, order := range orders {
		var user userModel.Users
		var product productModel.Products
		database.Database.Db.Find(&user, "id = ?", order.UserID)
		database.Database.Db.Find(&product, "id=?", order.ProductID)
		responseUser := userResponse.UserResponse{}.CreateUserResponse(user)
		responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
		responseOrder := orderResponse.OrderResponse{}.CreateOrderResponse(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func (orderHandler OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("order  not Found :id!=integer !!!")
	}
	var order orderModel.Orders
	orderFindError := orderService.OrderService{}.FindOrder(orderId, &order)
	if orderFindError != nil {
		return c.Status(400).JSON(orderFindError.Error())
	}
	var user userModel.Users
	var product productModel.Products
	database.Database.Db.First(&user, order.UserID)
	database.Database.Db.First(&product, order.ProductID)
	responseUser := userResponse.UserResponse{}.CreateUserResponse(user)
	responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
	responseOrder := orderResponse.OrderResponse{}.CreateOrderResponse(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func (orderHandler OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order orderModel.Orders
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user userModel.Users
	UserError := userService.UserService{}.FindUser(int(order.UserID), &user)
	if UserError != nil {
		return c.Status(400).JSON(UserError.Error())
	}

	var product productModel.Products
	ProductErr := productService.ProductService{}.FindProduct(int(order.ProductID), &product)
	if ProductErr != nil {
		return c.Status(400).JSON(ProductErr.Error())
	}
	database.Database.Db.Create(&order)
	responseUser := userResponse.UserResponse{}.CreateUserResponse(user)
	responseProduct := productResponse.ProductResponse{}.CreateProductResponse(product)
	responseOrder := orderResponse.OrderResponse{}.CreateOrderResponse(order, responseUser, responseProduct)
	data := map[string]interface{}{
		"success": "Order Created Successfully",
		"status":  201,
		"order":   responseOrder,
	}
	return c.Status(201).JSON(data)
}
