package order

import (
	"errors"
	"temporal-ecommerce/src/domain/entities"
)

func (s *OrderService) CreateOrder(order entities.Order) (entities.Order, error) {

	order.Status = entities.StatusPending

	order, err := s.orderRepository.Create(order)
	if err != nil {
		return entities.Order{}, err
	}

	product, err := s.productService.GetProduct(order.ProductID)
	if err != nil {
		return entities.Order{}, err
	}

	if product.Stock < uint(order.Quantity) {
		return entities.Order{}, errors.New("product stock is not enough")
	}

	order.Status = entities.StatusStockValidated
	order.Price = int(product.Price) * order.Quantity

	order, err = s.orderRepository.Update(order)
	if err != nil {
		return entities.Order{}, err
	}

	user, err := s.userService.GetUser(order.UserID)
	if err != nil {
		return entities.Order{}, err
	}

	if user.Wallet < order.Price {
		return entities.Order{}, errors.New("user balance is not enough")
	}

	order, err = s.orderRepository.Update(order)
	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}
