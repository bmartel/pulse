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
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		DbUsername, DbPassword, DbHost, DbPort, DbDatabase, DbSSL))
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
