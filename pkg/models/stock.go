package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Stock struct {
	LineNo   int
	Quantity int
	// TODO: Implement Product
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Stock{})
}