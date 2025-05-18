package handlers

import (
	orderservice "temporal-ecommerce/src/domain/order"
	productservice "temporal-ecommerce/src/domain/product"
	userservice "temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories"
	"temporal-ecommerce/src/web/handlers/order"
	"temporal-ecommerce/src/web/handlers/product"
	"temporal-ecommerce/src/web/handlers/user"

	"go.temporal.io/sdk/client"
	"gorm.io/gorm"
)

type HandlerContainer struct {
	UserHandler    *user.UserHandler
	ProductHandler *product.ProductHandler
	OrderHandler   *order.OrderHandler
}

func NewHandlerContainer(db *gorm.DB, temporalClient client.Client) *HandlerContainer {
	return &HandlerContainer{
		UserHandler:    user.NewUserHandler(userservice.NewUserService(repositories.NewUserRepository(db))),
		ProductHandler: product.NewProductHandler(productservice.NewProductService(repositories.NewProductRepository(db))),
		OrderHandler: order.NewOrderHandler(
			orderservice.NewOrderService(temporalClient),
		),
	}
}
