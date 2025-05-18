package product

import (
	"temporal-ecommerce/src/web/handlers/dtos"

	"github.com/gofiber/fiber/v2"
)

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var productDTO dtos.ProductDTO
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
