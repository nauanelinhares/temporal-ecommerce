package main

import (
	"temporal-ecommerce/internal/gormutils"

	"gorm.io/gorm"
)

// ProductMigration defines the migration
// Note: This struct and its methods are part of the 'main' package now.
type ProductMigration struct {
}

type Product struct {
	gormutils.BaseModel
	Name        string `gorm:"not null"`
	Description string
	Price       uint `gorm:"not null;check:price > 0"`
	Stock       uint `gorm:"not null;check:stock >= 0"`
}

// Name returns the name of this migration
func (m *ProductMigration) Name() string {
	return "001_product_table"
}

// Up creates the products table
func (m *ProductMigration) Up(db *gorm.DB) error {

	return db.AutoMigrate(&Product{})
}

// Down drops the products table
func (m *ProductMigration) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&Product{})
}

// init registers this migration. Since this file is now 'package main',
// it calls the Register function defined in migrations.go (also package main).
