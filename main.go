package main

import (
	"fmt"
	"log"
	"temporal-ecommerce/internal/config"
	"temporal-ecommerce/src/web/handlers"
	"temporal-ecommerce/src/web/handlers/health"

	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	cfg, err := config.LoadConfig(".")

	if err != nil {
		panic("failed to load config")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	setupRoutes(app, db, c)

	healthHandler := health.NewHealthHandler()
	healthHandler.Routes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App, db *gorm.DB, c client.Client) {

	handlerContainer := handlers.NewHandlerContainer(db, c)
	handlerContainer.UserHandler.Routes(app)
	handlerContainer.ProductHandler.Routes(app)
	handlerContainer.OrderHandler.Routes(app)
}
