package entities

import "github.com/google/uuid"

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Price     int
	Status    Status
}

type Status string

const (
	StatusPending        Status = "pending"
	StatusCompleted      Status = "completed"
	StatusCancelled      Status = "cancelled"
	StatusPaymentFailed  Status = "payment_failed"
	StatusPaid           Status = "paid"
	StatusShipped        Status = "shipped"
	StatusStockValidated Status = "stock_validated"
)
