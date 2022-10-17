package routes

import (
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterOrderRoutes(app *fiber.App) {
	app.Get("/api/v1/orders", controllers.GetAllOrders)
	app.Get("api/v1/order/:ID", controllers.GetOrderByID)
	app.Post("api/v1/order", controllers.CreateOrder)
	// app.Put("api/v1/order/:ID", controllers.UpdateOrder)
	app.Delete("api/v1/order/:ID", controllers.DeleteOrder)
}
