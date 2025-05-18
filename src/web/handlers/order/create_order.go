package order

import (
	"temporal-ecommerce/src/web/handlers/dtos"

	"github.com/gofiber/fiber/v2"
)

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	orderDTO := dtos.OrderDTO{}

	if err := c.BodyParser(&orderDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	order := orderDTO.ToDomain()

	order, err := h.orderService.CreateOrder(c.Context(), order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
