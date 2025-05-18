package main

import (
	"flag"
	"fmt"
	"temporal-ecommerce/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Migration interface {
	Name() string
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

var migrations []Migration

func main() {
	direction := flag.String("direction", "up", "direction of migration")
	flag.Parse()

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

	migrations = append(migrations, &ProductMigration{})

	for _, migration := range migrations {
		if *direction == "up" {
			err = migration.Up(db)
		} else {
			err = migration.Down(db)
		}
		if err != nil {
			panic("failed to migrate")
		}
	}

}
