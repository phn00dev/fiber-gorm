package orderResponse

import (
	"fiber-gorm/models/orderModel"
	"fiber-gorm/response/productResponse"
	"fiber-gorm/response/userResponse"
)

type OrderResponse struct {
	ID      uint
	User    userResponse.UserResponse
	Product productResponse.ProductResponse
}

func (orderResponse OrderResponse) CreateOrderResponse(order orderModel.Orders, user userResponse.UserResponse,
	product productResponse.ProductResponse) OrderResponse {
	return OrderResponse{
		ID:      order.ID,
		User:    user,
		Product: product,
	}
}
