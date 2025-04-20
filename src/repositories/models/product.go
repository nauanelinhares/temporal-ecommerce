package models

import (
	"temporal-ecommerce/internal/gormutils"
)

type Product struct {
	gormutils.BaseModel
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null;check:price > 0"`
	Stock       uint    `gorm:"not null;check:stock >= 0"`
}
