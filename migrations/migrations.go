package main

import (
	"flag"
	"log"
	"sort"

	// ASSUMPTION: You need to create these packages or adjust paths
	"temporal-ecommerce/internal/config"
	"temporal-ecommerce/internal/database"

	"gorm.io/gorm"
)

// Migration interface defines the methods each migration must implement
type Migration interface {
	Name() string
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

// migrations holds all registered migrations via init() functions
// in files like 001_product_table.go (also package main)
var migrations []Migration

// Register adds a migration to the migrations registry
// Called by init() in other package main files in this directory.
func Register(migration Migration) {
	migrations = append(migrations, migration)
}

// GetMigrations returns all registered migrations
// (Kept for potential future use, though main directly accesses the slice)
func GetMigrations() []Migration {
	return migrations
}

func main() {
	direction := flag.String("direction", "up", "Migration direction: 'up' or 'down'")
	flag.Parse()

	// Load configuration (assuming LoadConfig loads DB details)
	cfg, err := config.LoadConfig(".") // Adjust path to config if needed
	if err != nil {
		log.Fatalf("❌ Could not load config: %v", err)
	}

	// Connect to database
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	log.Println("🚀 Database connection established.")

	// --- TODO: Implement proper migration tracking (e.g., using a gorm_migrations table) ---

	// Sort migrations by name to ensure consistent order
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Name() < migrations[j].Name()
	})

	switch *direction {
	case "up":
		log.Println("🏃 Running migrations UP...")
		for _, m := range migrations { // Use the global slice directly
			// TODO: Check if migration already applied
			log.Printf("  Applying migration: %s", m.Name())
			if err := m.Up(db); err != nil {
				log.Fatalf("❌ Failed to apply migration %s: %v", m.Name(), err)
			}
			// TODO: Record migration as applied
			log.Printf("  ✅ Successfully applied migration: %s", m.Name())
		}
		log.Println("✨ All migrations applied successfully.")
	case "down":
		log.Println("🏃 Running migrations DOWN...")
		// Run down migrations in reverse order
		for i := len(migrations) - 1; i >= 0; i-- {
			m := migrations[i]
			// TODO: Check if migration was applied before reverting
			log.Printf("  Reverting migration: %s", m.Name())
			if err := m.Down(db); err != nil {
				log.Fatalf("❌ Failed to revert migration %s: %v", m.Name(), err)
			}
			// TODO: Remove migration record
			log.Printf("  ✅ Successfully reverted migration: %s", m.Name())
		}
		log.Println("✨ All migrations reverted successfully.")
	default:
		log.Fatalf("❌ Invalid direction: %s. Use 'up' or 'down'.", *direction)
	}
}
