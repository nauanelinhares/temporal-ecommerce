package handlers

import (
	productservice "temporal-ecommerce/src/domain/product"
	userservice "temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories"
	"temporal-ecommerce/src/web/handlers/product"
	"temporal-ecommerce/src/web/handlers/user"

	"gorm.io/gorm"
)

type HandlerContainer struct {
	UserHandler    *user.UserHandler
	ProductHandler *product.ProductHandler
}

func NewHandlerContainer(db *gorm.DB) *HandlerContainer {
	return &HandlerContainer{
		UserHandler:    user.NewUserHandler(userservice.NewUserService(repositories.NewUserRepository(db))),
		ProductHandler: product.NewProductHandler(productservice.NewProductService(repositories.NewProductRepository(db))),
	}
}
