package dtos

import (
	"temporal-ecommerce/src/domain/entities"

	"github.com/google/uuid"
)

type OrderDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (r *OrderDTO) ToDomain() entities.Order {
	return entities.Order{
		UserID:    uuid.MustParse(r.UserID),
		ProductID: uuid.MustParse(r.ProductID),
		Quantity:  r.Quantity,
	}
}
