package activities

import (
	"context"
	"errors"
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/domain/product"
	"temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories/interfaces"

	"github.com/google/uuid"
)

type Activities struct {
	OrderRepository interfaces.OrderRepository
	ProductService  *product.ProductService
	UserService     *user.UserService
}

func (a *Activities) CreateOrderActivity(ctx context.Context, order entities.Order) (entities.Order, error) {
	order.Status = entities.StatusPending
	return a.OrderRepository.Create(order)
}

func (a *Activities) GetProductActivity(ctx context.Context, productIDStr string) (entities.Product, error) {
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		return entities.Product{}, errors.New("invalid product ID format")
	}
	return a.ProductService.GetProduct(productID)
}

func (a *Activities) ValidateStockActivity(ctx context.Context, product entities.Product, order entities.Order) (entities.Order, error) {
	if product.Stock < uint(order.Quantity) {
		return entities.Order{}, errors.New("product stock is not enough")
	}
	order.Status = entities.StatusStockValidated
	order.Price = int(product.Price) * order.Quantity

	return order, nil
}

func (a *Activities) UpdateOrderActivity(ctx context.Context, order entities.Order) (entities.Order, error) {
	return a.OrderRepository.Update(order)
}

func (a *Activities) GetUserActivity(ctx context.Context, userIDStr string) (entities.User, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return entities.User{}, errors.New("invalid user ID format")
	}
	return a.UserService.GetUser(userID)
}

func (a *Activities) ValidateUserBalanceActivity(ctx context.Context, user entities.User, order entities.Order) (entities.Order, error) {
	if user.Wallet < order.Price {
		return entities.Order{}, errors.New("user balance is not enough")
	}
	order.Status = entities.StatusPaid
	return order, nil
}
