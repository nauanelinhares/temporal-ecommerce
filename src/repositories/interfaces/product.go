package interfaces

import "temporal-ecommerce/src/domain/entities"

type ProductRepository interface {
	Create(product entities.Product) error
}
