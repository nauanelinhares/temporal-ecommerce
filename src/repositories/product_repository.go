package repositories

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product entities.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product entities.Product) error {
	productModel := models.Product{}.FromDomain(product)
	return r.db.Create(&productModel).Error
}
