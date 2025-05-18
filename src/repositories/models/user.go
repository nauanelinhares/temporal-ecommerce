package models

import "temporal-ecommerce/internal/gormutils"

type User struct {
	gormutils.BaseModel
	Username string `gorm:"unique not null"`
	Email    string `gorm:"unique not null"`
	Wallet   int    `gorm:"default:0"`
}
