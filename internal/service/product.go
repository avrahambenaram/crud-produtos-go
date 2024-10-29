package service

import (
	"errors"
	"fmt"

	"github.com/avrahambenaram/crud-produtos-go/internal/entity"
)

type ProductService struct{}

func (c ProductService) GetAllProducts() []entity.Product {
	products := []entity.Product{}
	entity.DB.Find(&products)
	return products
}

func (c ProductService) GetProductById(ID uint) (entity.Product, error) {
	product := entity.Product{}
	entity.DB.Where("ID = ?", ID).Find(&product)
	if product.Description == "" && product.Price == 0 {
		return entity.Product{}, errors.New("Product not found")
	}
	return product, nil
}

func (c ProductService) InsertProduct(product entity.Product) (entity.Product, error) {
	result := entity.DB.Create(&product)
	if result.RowsAffected != 1 {
		return entity.Product{}, fmt.Errorf("An error occurred while creating product %s", product.Description)
	}
	return product, nil
}

func (c ProductService) DeleteProduct(ID uint) error {
	result := entity.DB.Delete(&entity.Product{}, ID)
	if result.RowsAffected != 1 {
		return errors.New("Product not found")
	}
	return nil
}

func (c ProductService) UpdateProduct(product entity.Product) (entity.Product, error) {
	_, err := c.GetProductById(product.ID)
	if err != nil {
		return entity.Product{}, nil
	}
	entity.DB.Save(&product)
	return product, nil
}
