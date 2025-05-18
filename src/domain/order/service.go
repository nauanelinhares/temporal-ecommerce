package order

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
)

type OrderService struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderService(orderRepository interfaces.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) CreateOrder(order entities.Order) (entities.Order, error) {
	return s.orderRepository.Create(order)
}
