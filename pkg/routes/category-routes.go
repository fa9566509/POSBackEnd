package routes

import "github.com/gofiber/fiber/v2"

func RegisterCategoryRoutes(app *fiber.App) {
	app.Get("api/v1/categories", controllers.GetAllCateories)
	app.Get("api/v1/category/:ID", controllers.GetCategoryById)
	app.Post("api/v1/category", controllers.CreateCategory)
	app.Put("api/v1/category/:ID", controllers.UpdateCategory)
	app.Delete("api/v1/category/:ID", controllers.DeleteCategory)
}