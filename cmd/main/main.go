package main

import (
	"github.com/fa9566509/POSBackEnd/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.RegisterCustomerRoutes(app)
	routes.RegisterOrderRoutes(app)
	app.Listen(":3000")
}
