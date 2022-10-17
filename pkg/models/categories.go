package models

import (
	"github.com/fa9566509/POSBackEnd/pkg/config"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	CategoryName string
}

func init() {
	d, _ := config.Connect()
	db = d
	db.AutoMigrate(&Categories{})
}

func GetAllCateories() (categories []Categories) {
	db.Find(&categories)
	return categories
}

func CreateCategory(category Categories) Categories {
	db.Create(&category)
	return category
}

func GetCategoryByID(ID int64) (category Categories) {
	db.Where("ID = ?", ID).Find(&category)
	return category
}

func DeleteCategory(ID int64) Categories {
	category := Categories{}
	db.Where("ID = ?", ID).Delete(&category)
	return category
}