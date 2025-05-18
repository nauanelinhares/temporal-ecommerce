package user

import (
	"temporal-ecommerce/src/web/handlers/dtos"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userDTO dtos.CreateUserRequest
	err := c.BodyParser(&userDTO)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := userDTO.ToDomain()
	user, err = h.userService.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(userDTO.FromDomain(user))
}
