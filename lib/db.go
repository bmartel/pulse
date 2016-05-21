package lib

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Sqlite Connection
func Sqlite() *gorm.DB {
	db, err := gorm.Open("sqlite3", DbURL)
	if err != nil {
		panic(err)
	}
	return db
}

// Postgres Connection
func Postgres() *gorm.DB {

	var connURL string

	if DbURL != "" {
		connURL = DbURL
	} else {
		connURL = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			DbUsername, DbPassword, DbHost, DbPort, DbDatabase, DbSSL)
	}

	db, err := gorm.Open("postgres", connURL)
	if err != nil {
		panic(err)
	}
	return db
}

// Mysql Connection
func Mysql() *gorm.DB {

	var connURL string

	if DbURL != "" {
		connURL = DbURL
	} else {
		connURL = fmt.Sprintf(
			"%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
			DbUsername, DbPassword, DbHost, DbPort, DbDatabase)
	}

	db, err := gorm.Open("mysql", connURL)
	if err != nil {
		panic(err)
	}
	return db
}
