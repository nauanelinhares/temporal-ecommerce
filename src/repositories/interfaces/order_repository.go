package interfaces

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order entities.Order) (entities.Order, error)
	Get(id uuid.UUID) (entities.Order, error)
	Update(order entities.Order) (entities.Order, error)
}
