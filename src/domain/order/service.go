package order

import (
	"go.temporal.io/sdk/client"
)

type OrderService struct {
	temporalClient client.Client
}

func NewOrderService(temporalClient client.Client) *OrderService {
	return &OrderService{temporalClient: temporalClient}
}
