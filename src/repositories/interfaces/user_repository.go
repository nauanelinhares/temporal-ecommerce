package interfaces

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	Get(id uuid.UUID) (entities.User, error)
	Update(user entities.User) (entities.User, error)
}
