package migrations

import (
	"gorm.io/gorm"
)

// Migration interface defines the methods each migration must implement
type Migration interface {
	Name() string
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

// migrations holds all registered migrations
var migrations []Migration

// Register adds a migration to the migrations registry
func Register(migration Migration) {
	migrations = append(migrations, migration)
}

// GetMigrations returns all registered migrations
func GetMigrations() []Migration {
	return migrations
}
