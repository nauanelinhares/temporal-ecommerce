package user

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

type UserService struct {
	userRepository interfaces.UserRepository
}

func (s *UserService) CreateUser(user entities.User) (entities.User, error) {
	return s.userRepository.Create(user)
}
