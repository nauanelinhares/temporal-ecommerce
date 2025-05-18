package repositories

import (
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/repositories/interfaces"
	"temporal-ecommerce/src/repositories/models"

	"github.com/google/uuid"
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

func (r *userRepository) Get(id uuid.UUID) (entities.User, error) {
	userModel := models.User{}
	err := r.db.First(&userModel, "id = ?", id).Error
	if err != nil {
		return entities.User{}, err
	}

	user := userModel.ToDomain()

	return user, nil
}

func (r *userRepository) Update(user entities.User) (entities.User, error) {
	userModel := models.User{}
	userModel.FromDomain(user)

	err := r.db.Save(&userModel).Error
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
