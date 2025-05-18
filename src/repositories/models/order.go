package models

import (
	"temporal-ecommerce/internal/gormutils"
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type Order struct {
	gormutils.BaseModel
	UserID    uuid.UUID       `gorm:"not null"`
	ProductID uuid.UUID       `gorm:"not null"`
	Quantity  int             `gorm:"not null;check:quantity > 0"`
	Status    entities.Status `gorm:"not null;default:pending"`
	Price     int             `gorm:"not null"`
}

func (o *Order) FromDomain(order entities.Order) {
	o.BaseModel = gormutils.BaseModel{ID: order.ID}
	o.UserID = order.UserID
	o.ProductID = order.ProductID
	o.Quantity = order.Quantity
	o.Status = order.Status
}

func (o *Order) ToDomain() entities.Order {
	return entities.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		ProductID: o.ProductID,
		Quantity:  o.Quantity,
		Status:    o.Status,
	}
}
