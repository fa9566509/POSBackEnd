package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("posbackend.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot initiate database")
	}
	return db, err
}