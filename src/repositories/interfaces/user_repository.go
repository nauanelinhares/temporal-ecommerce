package interfaces

import "temporal-ecommerce/src/domain/entities"

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
}
