package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName   string `json:"UserName"`
	Password   string `json:"Password"`
	Role       string `json:"Role"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	IdentityNo string `json:"IdentityNo"`
	Street     string `json:"Street"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
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
func (user Users) CreateUser() Users {
	db.Create(&user)
	return user
}

// GET
func GetUserByID(ID int64) (user Users) {
	db.Where("ID = ?", ID).Find(&user)
	return user
}

// DELETE
func DeleteUser(ID int64) (user Users) {
	db.Where("ID = ?", ID).Delete(&user)
	return user
}

// PUT
func (newUser Users)UpdateUser(ID int64) (oldUser Users) {
	db.Where("ID = ?", ID).Model(&oldUser).Updates(&newUser)
	return newUser
}
