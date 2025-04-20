package health

import "github.com/gofiber/fiber/v2"

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h HealthHandler) Routes(router fiber.Router) {
	router.Get("/ping", h.HealthCheck)
}

func (h HealthHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("pong :)")
}
