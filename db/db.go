package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Connect() *gorm.DB {
	conn, err := gorm.Open(viper.GetString("DATABASE_TYPE"), viper.GetString("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	// Run any db migrations
	Migrate(conn)

	return conn
}
