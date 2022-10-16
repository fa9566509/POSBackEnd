package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type PurchaseItems struct {
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