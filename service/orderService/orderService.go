package orderService

import (
	"errors"
	"fiber-gorm/database"
	"fiber-gorm/models/orderModel"
)

type OrderService struct{}

func (orderService OrderService) FindOrder(id int, order *orderModel.Orders) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order Not Found")
	}
	return nil
}
