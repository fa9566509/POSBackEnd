package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Purchases struct {
	gorm.Model
	PurchaseNumber int
	//TODO: Implement Purchase Date
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Purchases{})
}

// GET
func GetAllPurchases() (purchases []Purchases) {
	db.Find(&purchases)
	return purchases
}

// POST
func (purchase Purchases)CreatePurchase() Purchases {
	db.Create(&purchase)
	return purchase
}

func GetPurchaseByID(ID int64) (purchase Purchases) {
	db.Where("ID = ?", ID).Find(&purchase)
	return purchase
}

func DeletePurchase(ID int64) (purchase Purchases) {
	db.Where("ID = ?", ID).Delete(&purchase)
	return purchase
}