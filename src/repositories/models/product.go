package models

import (
	"temporal-ecommerce/internal/gormutils"
	"temporal-ecommerce/src/domain/entities"
)

type Product struct {
	gormutils.BaseModel
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null;check:price > 0"`
	Stock       uint    `gorm:"not null;check:stock >= 0"`
}

func (Product) FromDomain(product entities.Product) Product {
	return Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
