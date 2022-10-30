package routes

import (
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterStockRoutes(app *fiber.App) {
	app.Get("api/v1/stock", controllers.GetAllStocks)
	app.Get("api/v1/stock/:ID", controllers.GetStockById)
	app.Post("api/v1/stock", controllers.CreateStock)
	app.Put("api/v1/stock/:ID", controllers.UpdateStock)
	app.Delete("api/v1/stock/:ID", controllers.DeleteStock)
}