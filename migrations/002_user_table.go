package main

import (
	"temporal-ecommerce/internal/gormutils"

	"gorm.io/gorm"
)

type UserMigration struct {
}

type User struct {
	gormutils.BaseModel
	Username string `gorm:"unique not null"`
	Email    string `gorm:"unique not null"`
	Wallet   int    `gorm:"default:0"`
}

func (m *UserMigration) Name() string {
	return "002_user_table"
}

func (m *UserMigration) Up(db *gorm.DB) error {

	return db.AutoMigrate(&User{})
}

func (m *UserMigration) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&User{})
}
