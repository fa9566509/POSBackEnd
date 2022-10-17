package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	LineNo   int
	Quantity int
	// TODO: Implement Product
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Stock{})
}

// GET
func GetAllStocks() (stocks []Stock) {
	db.Find(&stocks)
	return stocks
}

// POST
func (stock Stock)CreateSock() Stock {
	db.Create(stock)
	return stock
}

func GetStockByID(ID int64) (stock Stock) {
	db.Where("ID = ?", ID).Find(&stock)
	return stock
}

func DeleteStock(ID int64) (stock Stock) {
	db.Where("ID = ?", ID).Delete(&stock)
	return stock
}