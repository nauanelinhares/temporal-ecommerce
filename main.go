package main

import (
	"log"
	"temporal-ecommerce/src/web/handlers/health"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	healthHandler := health.NewHealthHandler()
	healthHandler.Routes(app)

	log.Fatal(app.Listen(":3000"))
}
