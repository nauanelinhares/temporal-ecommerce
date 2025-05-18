package order

import (
	"temporal-ecommerce/src/domain/order"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService *order.OrderService
}

func NewOrderHandler(orderService *order.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Routes(router fiber.Router) {
	groupOrder := router.Group("/order")
	groupOrder.Post("/", h.CreateOrder)
}
