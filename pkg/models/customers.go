package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
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

func GetAllCustomers() (customers []Customers) {
	db.Find(&customers)
	return customers
}

func (customer Customers)CreateCustomer() Customers {
	db.Create(&customer)
	return customer
}

func GetCustomerByID(ID int64) (customer Customers) {
	db.Where("ID = ?", ID).Find(&customer)
	return customer
}

func DeleteCustomer(ID int64) Customers {
	customer := Customers{}
	db.Where("ID = ?", ID).Delete(&customer)
	return customer
}

func (newCustomer Customers)UpdateCustomer(ID int64) (oldCustomer Customers) {
	db.Where("ID = ?", ID).Model(&oldCustomer).Updates(&newCustomer)
	return newCustomer
}