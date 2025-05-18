package order

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/domain/product"
	"temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories/interfaces"
)

type OrderService struct {
	orderRepository interfaces.OrderRepository
	productService  *product.ProductService
	userService     *user.UserService
}

func NewOrderService(orderRepository interfaces.OrderRepository, productService *product.ProductService, userService *user.UserService) *OrderService {
	return &OrderService{orderRepository: orderRepository, productService: productService, userService: userService}
}

func (s *OrderService) CreateOrder(order entities.Order) (entities.Order, error) {

	order.Status = entities.StatusPending

	return s.orderRepository.Create(order)
}
