package routes

import (
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterStoreRoutes(app *fiber.App) {
	app.Get("api/v1/store", controllers.GetAllStore)
	app.Get("api/v1/store/:ID", controllers.GetStoreById)
	app.Post("api/v1/store", controllers.CreateStore)
	app.Put("api/v1/store/:ID", controllers.UpdateStore)
	app.Delete("api/v1/store/:ID", controllers.DeleteStore)
}