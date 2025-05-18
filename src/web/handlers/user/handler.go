package user

import (
	"temporal-ecommerce/src/domain/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Routes(router fiber.Router) {
	groupUser := router.Group("/user")
	groupUser.Post("/", h.CreateUser)
}
