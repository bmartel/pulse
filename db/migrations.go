package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
	// {
	// 	ID: "",
	// 	Migrate: func(tx *gorm.DB) error {
	// 		return tx.AutoMigrate().Error
	// 	},
	// 	Rollback: func(tx *gorm.DB) error {
	// 		return tx.DropTable("").Error
	// 	},
	// },
	}
}

func Migrate(db *gorm.DB) {
	migrate := migrations()
	if len(migrate) < 1 {
		log.Println("nothing to migrate")
		return
	}
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrate)

	if err := m.Migrate(); err != nil {
		log.Printf("could not migrate db: %v", err)
	}
}
