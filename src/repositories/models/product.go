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

func (p *Product) FromDomain(product entities.Product) {
	if product.ID == uuid.Nil {
		product.ID = uuid.New()
	}
	p.BaseModel = gormutils.BaseModel{ID: product.ID}
	p.Name = product.Name
	p.Description = product.Description
	p.Price = product.Price
	p.Stock = product.Stock
}
