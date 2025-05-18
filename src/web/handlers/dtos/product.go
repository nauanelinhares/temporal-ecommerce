package dtos

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type ProductDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Stock       uint      `json:"stock"`
}

func (p *ProductDTO) ToDomain() entities.Product {
	return entities.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
	}
}

func (p *ProductDTO) FromDomain(product entities.Product) ProductDTO {
	return ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
