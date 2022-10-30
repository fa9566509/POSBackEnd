package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	OrderStatus string  `json:"OrderStatus"`
	TotalPrice  float64 `json:"TotalPrice"`
	Discount    float32 `json:"Discount"`
	Quantity    int     `json:"Quantity"`
	LineNo      int     `json:"LineNo"`
	OrdersID    uint    `json:"OrderID"`
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
func (orderItem OrderItems) CreateOrderItem() OrderItems {
	db.Create(&orderItem)
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

func (orderItem OrderItems) UpdateOrderItem(ID int64) (oldOrderItem OrderItems) {
	db.Where("ID = ?", ID).Model(&oldOrderItem).Updates(&orderItem)
	return orderItem
}
