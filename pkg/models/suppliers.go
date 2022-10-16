package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Suppliers struct {
	FirstName    string
	LastName     string
	Company      string
	Street       string
	City         string
	Country      string
	SupplierCode string
	Email        string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Suppliers{})
}