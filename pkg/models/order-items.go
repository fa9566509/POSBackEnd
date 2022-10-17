package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	OrderStatus string
	TotalPrice  float64
	Discount    float32
	Quantity    int
	LineNo      int
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&OrderItems{})
}

// GET
func GetAllOrdersItems() (orderItems []OrderItems) {
	db.Find(&orderItems)
	return orderItems
}

// POST
func (orderItem OrderItems)CreateOrderItem() OrderItems {
	db.Create(orderItem)
	return orderItem
}

func GetOrderItemByID(ID int64) (orderItem OrderItems) {
	db.Where("ID = ?", ID).Find(&orderItem)
	return orderItem
}

func DeleteOrderItem(ID int64) (orderItem OrderItems) {
	db.Where("ID = ?", ID).Delete(&orderItem)
	return orderItem
}