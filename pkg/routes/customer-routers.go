package routes

import (
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)


func RegisterCustomerRoutes(app *fiber.App) {
	app.Get("api/v1/customers", controllers.GetCustomers)
	app.Get("api/v1/customer/:ID", controllers.GetCustomerByID)
	app.Post("api/v1/customer", controllers.CreateCustomer)
	app.Put("api/v1/customer/:ID", controllers.UpdateCustomer)
	app.Delete("api/v1/customer/:ID", controllers.DeleteCustomer)
}