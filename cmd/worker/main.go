package main

import (
	"fmt"
	"log"
	"temporal-ecommerce/internal/config"
	"temporal-ecommerce/src/domain/product"
	"temporal-ecommerce/src/domain/user"
	"temporal-ecommerce/src/repositories"
	"temporal-ecommerce/src/temporal/order/activities"
	"temporal-ecommerce/src/temporal/order/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	cfg, err := config.LoadConfig(".")

	if err != nil {
		panic("failed to load config")
	}

	defer c.Close()

	w := worker.New(c, "ORDER_TASK_QUEUE", worker.Options{})

	w.RegisterWorkflow(workflows.CreateOrderWorkflow)

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

	orderRepository := repositories.NewOrderRepository(db)
	productRepository := repositories.NewProductRepository(db)
	userRepository := repositories.NewUserRepository(db)
	productService := product.NewProductService(productRepository)
	userService := user.NewUserService(userRepository)

	acts := &activities.Activities{
		OrderRepository: orderRepository,
		ProductService:  productService,
		UserService:     userService,
	}
	w.RegisterActivity(acts.CreateOrderActivity)
	w.RegisterActivity(acts.GetProductActivity)
	w.RegisterActivity(acts.ValidateStockActivity)
	w.RegisterActivity(acts.UpdateOrderActivity)
	w.RegisterActivity(acts.GetUserActivity)
	w.RegisterActivity(acts.ValidateUserBalanceActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
