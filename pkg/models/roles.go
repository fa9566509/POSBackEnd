package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Sock struct {
	gorm.Model
	RoleName string
	//TODO: Implement role value
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Sock{})
}

// GET
func GetAllRoles() (roles []Sock) {
	db.Find(&roles)
	return roles
}

// POST
func (role Sock) CreateRole() Sock {
	db.Create(&role)
	return role
}

func GetRoleByID(ID int64) (role Sock) {
	db.Where("ID = ?", ID).Find(&role)
	return role
}

func DeleteRole(ID int64) (role Sock) {
	db.Where("ID = ?", ID).Delete(&role)
	return role
}

func (role Sock) UpdateRole(ID int64) (oldRole Sock) {
	db.Where("ID = ?", ID).Model(&oldRole).Updates(&role)
	return role
}
