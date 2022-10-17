package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type PurchaseItems struct {
	gorm.Model
	PurchaseStaus string
	TotalPrice    float64
	Discount      float32
	Quantity      int
	LineNo        int
	// TODO: Implement Relation
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&PurchaseItems{})
}

// GET
func GetAllPurchaseItems() (purchaseItems []PurchaseItems) {
	db.Find(&purchaseItems)
	return purchaseItems
}

// POST
func (purchaseItem PurchaseItems)CreatePurchaseItem() PurchaseItems {
	db.Create(purchaseItem)
	return purchaseItem
}

func GetPurchaseItemByID(ID int64) (purchaseItem PurchaseItems) {
	db.Where("ID = ?", ID).Find(&purchaseItem)
	return purchaseItem
}

func DeletePurchaseItem(ID int64) (purchaseItem PurchaseItems) {
	db.Where("ID = ?", ID).Delete(&purchaseItem)
	return purchaseItem
}