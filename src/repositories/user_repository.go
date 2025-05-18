package repositories

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
	"temporal-ecommerce/src/repositories/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entities.User) (entities.User, error) {
	userModel := models.User{}
	userModel.FromDomain(user)

	err := r.db.Create(&userModel).Error
	if err != nil {
		return entities.User{}, err
	}

	user.ID = userModel.ID

	return user, nil
}
