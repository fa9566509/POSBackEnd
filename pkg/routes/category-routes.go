package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
)

func RegisterCategoryRoutes(app *fiber.App) {
	app.Get("api/v1/categories", controllers.GetAllCategories)
	app.Get("api/v1/category/:ID", controllers.GetCategoryByID)
	app.Post("api/v1/category", controllers.CreateCategory)
	app.Put("api/v1/category/:ID", controllers.UpdateCategory)
	app.Delete("api/v1/category/:ID", controllers.DeleteCategory)
}