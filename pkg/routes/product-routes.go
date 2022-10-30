package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fa9566509/POSBackEnd/pkg/controllers"
)

func RegisterProductRoutes(app *fiber.App) {
	app.Get("api/v1/products", controllers.GetAllProducts)
	app.Get("api/v1/product/:ID", controllers.GetProductById)
	app.Post("api/v1/product", controllers.CreateProduct)
	app.Put("api/v1/product/:ID", controllers.UpdateProduct)
	app.Delete("api/v1/product/:ID", controllers.DeleteProduct)
}
