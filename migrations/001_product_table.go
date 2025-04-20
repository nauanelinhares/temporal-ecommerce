package migrations

import (
	"temporal-ecommerce/src/repositories/models"

	"gorm.io/gorm"
)

// Up creates the products table
func (m *ProductMigration) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})
}

// Down drops the products table
func (m *ProductMigration) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Product{})
}

// ProductMigration defines the migration
type ProductMigration struct{}

// Name returns the name of this migration
func (m *ProductMigration) Name() string {
	return "001_product_table"
}

// Register registers this migration
func init() {
	Register(&ProductMigration{})
}
