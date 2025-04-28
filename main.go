package main

import (
	"fmt"
	"log"
	"temporal-ecommerce/internal/config"
	productservice "temporal-ecommerce/src/domain/product"
	"temporal-ecommerce/src/repositories"
	"temporal-ecommerce/src/web/handlers/health"
	"temporal-ecommerce/src/web/handlers/product"

	"github.com/gofiber/fiber/v2"
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

	setupRoutes(app, db)

	healthHandler := health.NewHealthHandler()
	healthHandler.Routes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App, db *gorm.DB) {

	productRepository := repositories.NewProductRepository(db)
	productService := productservice.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)
	productHandler.Routes(app)
}
