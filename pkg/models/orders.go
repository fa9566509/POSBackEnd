package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Orders struct {
	gorm.Model
	OrderNumber int          `json:"OrderNumber"`
	OrderItem   []OrderItems `gorm:"foreignKey:OrdersID" json:"OrderItems"`
}

func init() {
	con, _ := config.Connect()
	db = con
	con.AutoMigrate(&Orders{})
}

func GetAllOrders() (order []Orders) {
	db.Preload("OrderItem").Find(&order)
	return order
}

func (order Orders) CreateOrder() Orders {
	db.Create(&order)
	db.Save(&order)
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


func (order Orders) UpdateOrder(ID int64) (oldOrder Orders) {
	db.Where("ID = ?", ID).Model(&oldOrder).Updates(&order)
	return order
}