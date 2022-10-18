package main

import (
	"github.com/fa9566509/POSBackEnd/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	utils.RegisterAllRoutes(app)
	app.Listen(":9010")
}
