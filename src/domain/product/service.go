package product

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"

	"github.com/google/uuid"
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

func (s *ProductService) GetProduct(id uuid.UUID) (entities.Product, error) {
	product, err := s.productRepository.Get(id)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(product entities.Product) (entities.Product, error) {
	product, err := s.productRepository.Update(product)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}
