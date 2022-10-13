package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	con, _ := config.Connect()
	db = con
	con.AutoMigrate(&Orders{})
}

type Orders struct {
	OrderNumber int
	//TODO: Add Order Date
}

func GetAllOrders() (orders []Orders) {
	db.Find(&orders)
	return orders
}