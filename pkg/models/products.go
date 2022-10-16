package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Products struct {
	LineNo    int
	Barcode   string
	ALU       string
	ProdName  string
	BuyPrice  float32
	SalePrice float32
	// TODO: Implement Supplier relationship many to many
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Products{})
}