package productService

import (
	"errors"
	"fiber-gorm/database"
	"fiber-gorm/models/productModel"
)

type ProductService struct{}

func (productService ProductService) FindProduct(id int, product *productModel.Products) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("Product not found!!!")
	}
	return nil
}
