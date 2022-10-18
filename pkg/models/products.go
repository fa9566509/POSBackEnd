package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	LineNo    int
	Barcode   string
	ALU       string
	ProdName  string
	BuyPrice  float32
	SalePrice float32
	// TODO: Implement Supplier relationship many to many
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Products{})
}

// GET
func GetAllProducts() (products []Products) {
	db.Find(&products)
	return products
}

// POST
func (product Products)CreateProduct() Products {
	db.Create(&product)
	return product
}

func GetProductByID(ID int64) (product Products) {
	db.Where("ID = ?", ID).Find(&product)
	return product
}

func DeleteProduct(ID int64) (product Products) {
	db.Where("ID = ?", ID).Delete(&product)
	return product
}