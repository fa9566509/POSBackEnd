package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Stores struct {
	gorm.Model
	StoreName string
	Street    string
	City      string
	State     string
	Country   string
	PhoneNo   string
	Email     string
	Website   string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Stores{})
}

func GetAllStores() (stores []Stores) {
	db.Find(&stores)
	return stores
}

func (store Stores)CreateStore() Stores {
	db.Create(&store)
	return store
}

func GetStoreByID(ID int64) (store Stores) {
	db.Where("ID = ?", ID).Find(&store)
	return store
}

func DeleteStore(ID int64) (store Stores) {
	db.Where("ID = ?", ID).Find(&store)
	return store
}