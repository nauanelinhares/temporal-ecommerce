package gormutils

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID != uuid.Nil {
		return nil
	}
	base.ID = uuid.New()
	return nil
}
