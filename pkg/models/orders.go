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
	gorm.Model
	OrderNumber int `json:"OrderNumber"`
	//TODO: Add Order Date
}

func GetAllOrders() (orders []Orders) {
	db.Find(&orders)
	return orders
}

func (order Orders) CreateOrder() Orders {
	db.Create(&order)
	return order
}

func GetOrderbyID(ID int64) Orders {
	order := Orders{}
	db.Where("ID = ?", ID).Find(&order)
	return order
}

func DeleteOrder(ID int64) (order Orders) {
	db.Where("ID = ?", ID).Delete(&order)
	return order
}