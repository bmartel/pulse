package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Sqlite ...
func Sqlite() *gorm.DB {
	db, err := gorm.Open("sqlite3", DbDatabase)
	if err != nil {
		panic(err)
	}
	return db
}

// Postgres ...
func Postgres() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s:%s user=%s dbname=%s sslmode=%s password=%s",
		DbHost, DbPort, DbUsername, DbDatabase, DbSSL, DbPassword))
	if err != nil {
		panic(err)
	}
	return db
}

// Mysql ...
func Mysql() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
		DbUsername, DbPassword, DbHost, DbPort, DbDatabase))
	if err != nil {
		panic(err)
	}
	return db
}
