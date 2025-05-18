package repositories

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
	"temporal-ecommerce/src/repositories/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product entities.Product) (entities.Product, error) {
	productModel := models.Product{}
	productModel.FromDomain(product)

	err := r.db.Create(&productModel).Error
	if err != nil {
		return entities.Product{}, err
	}

	product.ID = productModel.ID

	return product, nil
}

func (r *productRepository) Get(id uuid.UUID) (entities.Product, error) {
	productModel := models.Product{}
	err := r.db.First(&productModel, "id = ?", id).Error
	if err != nil {
		return entities.Product{}, err
	}

	product := productModel.ToDomain()

	return product, nil
}

func (r *productRepository) Update(product entities.Product) (entities.Product, error) {
	productModel := models.Product{}
	productModel.FromDomain(product)

	err := r.db.Save(&productModel).Error
	if err != nil {
		return entities.Product{}, err
	}

	return product, nil
}
