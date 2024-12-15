package productResponse

import "fiber-gorm/models/productModel"

type ProductResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	TotalCount  int    `json:"total_count"`
}

type ProductUpdate struct {
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	TotalCount  int    `json:"total_count"`
}

func (productResponse ProductResponse) CreateProductResponse(product productModel.Products) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		Price:       product.Price,
		TotalCount:  product.TotalCount,
	}
}
