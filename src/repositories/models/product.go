package models

import (
	"temporal-ecommerce/internal/gormutils"
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type Product struct {
	gormutils.BaseModel
	Name        string `gorm:"not null"`
	Description string
	Price       uint `gorm:"not null;check:price > 0"`
	Stock       uint `gorm:"not null;check:stock >= 0"`
}

func (Product) FromDomain(product entities.Product) Product {
	if product.ID == uuid.Nil {
		product.ID = uuid.New()
	}
	return Product{
		BaseModel:   gormutils.BaseModel{ID: product.ID},
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
