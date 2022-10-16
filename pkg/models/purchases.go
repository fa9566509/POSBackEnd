package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Purchases struct {
	PurchaseNumber int
	//TODO: Implement Purchase Date
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Purchases{})
}