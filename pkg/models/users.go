package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Users struct {
	UserName   string
	Password   string
	Role       string
	FirstName  string
	LastName   string
	IdentityNo string
	Street     string
	City       string
	State      string
	Country    string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Users{})
}