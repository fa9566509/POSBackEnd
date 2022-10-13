package routes

import "github.com/gofiber/fiber/v2"

func RegisterPurchaseItemRoutes(app *fiber.App) {
	app.Get("api/v1/purchaseitems", controllers.GetAllPurchaseItems)
	app.Get("api/v1/purchaseitems/:ID", controllers.GetPurchaseItemById)
	app.Post("api/v1/purchaseitems", controllers.CreatePurchaseItem)
	app.Put("api/v1/purchaseitems/:ID", controllers.UpdatePurchaseItem)
	app.Delete("api/v1/purchaseitems/:ID", controllers.DeletePurchaseItem)
}