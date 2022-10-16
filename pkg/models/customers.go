package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	//"gorm.io/gorm"
)

type Customers struct {
	FirstName string
	LastName  string
	CardNo    string
	Email     string
	PhoneNo   string
	Street    string
	City      string
	State     string
	Country   string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Customers{})
}