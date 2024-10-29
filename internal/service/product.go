package service

import "github.com/avrahambenaram/crud-produtos-go/internal/entity"

type ProductService struct{}

func (c ProductService) GetAllProducts() []entity.Product {
	products := []entity.Product{}
	entity.DB.Find(&products)
	return products
}
