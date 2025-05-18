package dtos

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Wallet   int       `json:"wallet"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u *CreateUserRequest) ToDomain() entities.User {
	return entities.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Wallet:   u.Wallet,
	}
}

func (u *CreateUserRequest) FromDomain(user entities.User) CreateUserRequest {
	return CreateUserRequest{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Wallet:   user.Wallet,
	}
}
