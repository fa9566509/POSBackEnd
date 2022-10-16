package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type OrderItems struct {
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