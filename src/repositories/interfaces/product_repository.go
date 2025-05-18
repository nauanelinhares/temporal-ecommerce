package interfaces

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product entities.Product) (entities.Product, error)
	Get(id uuid.UUID) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
}
