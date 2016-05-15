package config

import "github.com/jinzhu/gorm"

// AutoMigrate ... Migrations to run on application boot
func AutoMigrate(db *gorm.DB) {
	// Preload any migrations that need to exist
}
