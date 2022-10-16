package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Stores struct {
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