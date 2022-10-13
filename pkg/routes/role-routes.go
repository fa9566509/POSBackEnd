package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoleRoutes(app *fiber.App) {
	app.Get("api/v1/roles", controllers.GetAllRoles)
	app.Get("api/v1/role/:ID", controllers.GetRoleById)
	app.Post("api/v1/role", controllers.CreateRole)
	app.Put("api/v1/role/:ID", controllers.UpdateRole)
	app.Delete("api/v1/role/:ID", controllers.DeleteRole)
}