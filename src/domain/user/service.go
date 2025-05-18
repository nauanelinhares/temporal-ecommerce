package user

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"

	"github.com/google/uuid"
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

func (s *UserService) GetUser(id uuid.UUID) (entities.User, error) {
	return s.userRepository.Get(id)
}

func (s *UserService) UpdateUser(user entities.User) (entities.User, error) {
	return s.userRepository.Update(user)
}
