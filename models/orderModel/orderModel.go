package orderModel

import (
	"fiber-gorm/models/productModel"
	"fiber-gorm/models/userModel"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID    uint                  `json:"user_id"`
	User      userModel.Users       `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint                  `json:"product_id"`
	Product   productModel.Products `json:"product" gorm:"foreignKey:ProductID"`
}
