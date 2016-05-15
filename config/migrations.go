package config

import "github.com/jinzhu/gorm"

func AutoMigrate(db *gorm.DB) {
	// Preload any migrations that need to exist
}
