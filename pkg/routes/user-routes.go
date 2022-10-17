package routes

import (
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	app.Get("api/v1/users", controllers.GetAllUsers)
	app.Get("api/v1/user/:ID", controllers.GetUserByID)
	// app.Post("api/v1/user", controllers.CreateUser)
	// app.Put("api/v1/user/:ID", controllers.UpdateUser)
	// app.Delete("api/v1/user/:ID", controllers.DeleteUser)
}
