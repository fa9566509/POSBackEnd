package models

import "github.com/fa9566509/POSBackEnd/pkg/config"

type Roles struct {
	RoleName string
	//TODO: Implement role value
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Roles{})
}