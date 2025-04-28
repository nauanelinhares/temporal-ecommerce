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

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var productDTO ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product := productDTO.ToDomain()
	product, err := h.productService.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}
