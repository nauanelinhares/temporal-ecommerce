package entities

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Price     int
	Status    string
}
