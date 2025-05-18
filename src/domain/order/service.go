package order

import (
	"temporal-ecommerce/src/domain/product"
	"temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories/interfaces"

	"go.temporal.io/sdk/client"
)

type OrderService struct {
	orderRepository interfaces.OrderRepository
	productService  *product.ProductService
	userService     *user.UserService
	temporalClient  client.Client
}

func NewOrderService(orderRepository interfaces.OrderRepository, productService *product.ProductService, userService *user.UserService, temporalClient client.Client) *OrderService {
	return &OrderService{orderRepository: orderRepository, productService: productService, userService: userService, temporalClient: temporalClient}
}
