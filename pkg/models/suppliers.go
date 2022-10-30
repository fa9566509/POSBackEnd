package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Suppliers struct {
	gorm.Model
	FirstName    string
	LastName     string
	Company      string
	Street       string
	City         string
	Country      string
	SupplierCode string
	Email        string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Suppliers{})
}

func GetAllSuppliers() (suppliers []Suppliers) {
	db.Find(&suppliers)
	return suppliers
}

func (supplier Suppliers)CreateSupplier() Suppliers {
	db.Create(&supplier)
	return supplier
}

func GetSupplierByID(ID int64) (supplier Suppliers) {
	db.Where("ID = ?", ID).Find(&supplier)
	return supplier
}

func DeleteSupplier(ID int64) (supplier Suppliers) {
	db.Where("ID = ?", ID).Delete(&supplier)
	return supplier
}

func (supplier Suppliers) UpdateSupplier(ID int64) (oldSupplier Suppliers) {
	db.Where("ID = ?", ID).Model(&oldSupplier).Updates(&supplier)
	return supplier
}