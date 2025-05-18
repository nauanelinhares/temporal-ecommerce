package product

import (
	"temporal-ecommerce/src/domain/product"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService *product.ProductService
}

func NewProductHandler(productService *product.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) Routes(router fiber.Router) {
	groupProduct := router.Group("/product")
	groupProduct.Post("/", h.CreateProduct)
}
