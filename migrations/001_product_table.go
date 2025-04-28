package main

import (
	"temporal-ecommerce/src/repositories/models"

	"gorm.io/gorm"
)

// ProductMigration defines the migration
// Note: This struct and its methods are part of the 'main' package now.
type ProductMigration struct {
}

// Name returns the name of this migration
func (m *ProductMigration) Name() string {
	return "001_product_table"
}

// Up creates the products table
func (m *ProductMigration) Up(db *gorm.DB) error {

	return db.AutoMigrate(&models.Product{})
}

// Down drops the products table
func (m *ProductMigration) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Product{})
}

// init registers this migration. Since this file is now 'package main',
// it calls the Register function defined in migrations.go (also package main).
