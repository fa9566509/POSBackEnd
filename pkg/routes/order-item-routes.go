package routes

import "github.com/gofiber/fiber/v2"

func RegisterOrderItemRoutes(app *fiber.App) {
	app.Get("api/v1/orderitems", controllers.GetOrderItems)
	app.Get("api/v1/orderitem/:ID", controllers.GetOrderItemByID)
	app.Post("api/v1/orderitem", controllers.CreateOrderItem)
	app.Put("api/v1/orderitem/:ID", controllers.UpdateOrderItem)
	app.Delete("api/v1/orderitem/:ID", controllers.DeleteOrderItem)
}