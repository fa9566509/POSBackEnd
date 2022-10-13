package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	// TODO: Need to implement GetCustomers wih Gorm.io

	return c.SendString("Hello Customers")
}

func GetCustomerByID(c *fiber.Ctx) error {
	// TODO: Need to Implement GetCustomerByID with Gorm.io

	return c.SendString("GetCustomerByID")
}
