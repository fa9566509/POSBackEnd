package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
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

// GET
func GetAllUsers() (users []Users) {
	db.Find(&users)
	return users
}

// POST
func (user Users)CreateUser() Users {
	db.Create(user)
	return user
}

func GetUserByID(ID int64) (user Users) {
	db.Where("ID = ?", ID).Find(&user)
	return user
}

func DeleteUser(ID int64) (user Users) {
	db.Where("ID = ?", ID).Delete(&user)
	return user
}