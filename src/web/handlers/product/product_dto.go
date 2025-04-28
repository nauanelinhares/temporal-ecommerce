package product

import (
	"temporal-ecommerce/src/domain/entities"
)

type ProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
}

func (p *ProductDTO) ToDomain() entities.Product {
	return entities.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
	}
}

func (p *ProductDTO) FromDomain(product entities.Product) ProductDTO {
	return ProductDTO{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
