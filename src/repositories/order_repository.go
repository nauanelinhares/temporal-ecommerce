package repositories

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
	"temporal-ecommerce/src/repositories/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func (r *orderRepository) NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order entities.Order) (entities.Order, error) {
	orderModel := models.Order{}
	orderModel.FromDomain(order)

	err := r.db.Create(&orderModel).Error
	if err != nil {
		return entities.Order{}, err
	}

	order.ID = orderModel.ID

	return order, nil
}

func (r *orderRepository) Get(id uuid.UUID) (entities.Order, error) {
	orderModel := models.Order{}
	err := r.db.First(&orderModel, "id = ?", id).Error
	if err != nil {
		return entities.Order{}, err
	}

	order := orderModel.ToDomain()

	return order, nil
}

func (r *orderRepository) Update(order entities.Order) (entities.Order, error) {
	orderModel := models.Order{}
	orderModel.FromDomain(order)

	err := r.db.Save(&orderModel).Error
	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}
