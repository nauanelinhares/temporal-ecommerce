package product

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
)

type ProductService struct {
	productRepository interfaces.ProductRepository
}

func NewProductService(productRepository interfaces.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(product entities.Product) (entities.Product, error) {
	product, err := s.productRepository.Create(product)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}
