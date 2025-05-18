package entities

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       uint
	Stock       uint
}
