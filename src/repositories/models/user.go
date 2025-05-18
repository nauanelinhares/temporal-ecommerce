package models

import (
	"temporal-ecommerce/internal/gormutils"
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type User struct {
	gormutils.BaseModel
	Username string `gorm:"unique not null"`
	Email    string `gorm:"unique not null"`
	Wallet   int    `gorm:"default:0"`
}

func (u *User) FromDomain(user entities.User) {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	u.BaseModel = gormutils.BaseModel{ID: user.ID}
	u.Username = user.Username
	u.Email = user.Email
	u.Wallet = user.Wallet
}

func (u *User) ToDomain() entities.User {
	return entities.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Wallet:   u.Wallet,
	}
}
