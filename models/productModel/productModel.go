package productModel

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	TotalCount  int    `json:"total_count"`
}
