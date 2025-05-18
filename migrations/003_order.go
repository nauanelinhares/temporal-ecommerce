package main

import (
	"temporal-ecommerce/internal/gormutils"
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderMigration struct{}

type Order struct {
	gormutils.BaseModel
	UserID    uuid.UUID       `gorm:"not null"`
	ProductID uuid.UUID       `gorm:"not null"`
	Quantity  int             `gorm:"not null;check:quantity > 0"`
	Status    entities.Status `gorm:"not null;default:pending"`
}

func (m *OrderMigration) Name() string {
	return "003_order_table"
}

func (m *OrderMigration) Up(db *gorm.DB) error {

	return db.AutoMigrate(&Order{})
}

func (m *OrderMigration) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&Order{})
}
