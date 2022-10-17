package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	RoleName string
	//TODO: Implement role value
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Roles{})
}

// GET
func GetAllRoles() (roles []Roles) {
	db.Find(&roles)
	return roles
}

// POST
func (role Roles)CreateRole() Roles {
	db.Create(role)
	return role
}

func GetRoleByID(ID int64) (role Roles) {
	db.Where("ID = ?", ID).Find(&role)
	return role
}

func DeleteRole(ID int64) (role Roles) {
	db.Where("ID = ?", ID).Delete(&role)
	return role
}