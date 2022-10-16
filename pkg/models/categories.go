package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Categories struct {
	CategoryName string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Categories{})
}